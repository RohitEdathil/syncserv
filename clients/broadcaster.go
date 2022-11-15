package clients

import (
	"log"
	"sync"
	"time"

	e "syncserv/error_handling"

	"syncserv/util"

	"github.com/gorilla/websocket"
)

// A Broadcaster is a websocket connection that can send messages to all its listeners
type Broadcaster struct {
	Id         string
	Secret     string
	Text       string
	Connection *websocket.Conn
	Lock       *sync.Mutex
	Listeners  map[int]Listener
	count      int
	GreenCount int
	LastSeen   time.Time

	ConnectedHandler    func(broadcaster *Broadcaster)
	MessageHandler      func(broadcaster *Broadcaster, message *util.Message)
	DisconnectedHandler func(broadcaster *Broadcaster)
}

// Starts listening to messages from the broadcaster
func (broadcaster *Broadcaster) StartListening(conn *websocket.Conn) {

	// Assign connection
	broadcaster.Lock.Lock()
	broadcaster.Connection = conn
	broadcaster.Lock.Unlock()

	broadcaster.ConnectedHandler(broadcaster)

	// Listen loop
	for {
		// Readig and parsing message
		message := util.Message{}
		err := broadcaster.Connection.ReadJSON(&message)

		// Error handling
		if err != nil {
			e.PanicWS(broadcaster.Connection, err.Error())
			broadcaster.Connection.Close()
			break
		}

		// Handling message
		log.Printf("Message received from %s", broadcaster.Connection.RemoteAddr())
		broadcaster.MessageHandler(broadcaster, &message)
	}

	// Broadcast disconnection
	log.Printf("Disconnected")

	broadcaster.DisconnectedHandler(broadcaster)

	broadcaster.Lock.Lock()
	broadcaster.Connection = nil
	broadcaster.Lock.Unlock()

}

// Adds a listener to the broadcaster
func (broadcaster *Broadcaster) AddListener(listener *Listener) {
	// Entry
	broadcaster.Lock.Lock()
	listener.Lock.Lock()

	// Critical
	listener.id = broadcaster.count
	broadcaster.Listeners[broadcaster.count] = *listener
	broadcaster.count++

	// Exit
	broadcaster.Lock.Unlock()
	listener.Lock.Unlock()
}

// Removes a listener from the broadcaster
func (broadcaster *Broadcaster) RemoveListener(listener *Listener) {
	broadcaster.Lock.Lock()
	delete(broadcaster.Listeners, listener.id)
	broadcaster.Lock.Unlock()
}
