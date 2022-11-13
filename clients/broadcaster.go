package clients

import (
	"log"
	"sync"

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

	Handler func(broadcaster *Broadcaster, message *util.Message)
}

// Starts listening to messages from the broadcaster
func (broadcaster *Broadcaster) StartListening(conn *websocket.Conn) {

	// Assign connection
	broadcaster.Lock.Lock()
	broadcaster.Connection = conn
	broadcaster.Connection.WriteJSON(util.Message{
		Type: "code-state",
		Data: broadcaster.Text,
	})
	broadcaster.Lock.Unlock()

	// Listen loop
	for {
		// Readig and parsing message
		message := util.Message{}
		err := broadcaster.Connection.ReadJSON(&message)

		// Error handling
		if err != nil {
			e.PanicWS(*broadcaster.Connection, err.Error())
			broadcaster.Connection.Close()
			break
		}

		// Handling message
		log.Printf("Message received from %s : %s", broadcaster.Connection.RemoteAddr(), message)
		broadcaster.Handler(broadcaster, &message)
	}

	// Broadcast disconnection
	log.Printf("Disconnected")

	broadcaster.Lock.Lock()
	broadcaster.Connection = nil
	broadcaster.Lock.Unlock()

	// Remove broadcaster from list if no more listeners
	if len(broadcaster.Listeners) == 0 {
		log.Printf("No more listeners, closing broadcaster")
		ClientIndexInstance.Delete(broadcaster.Id)
	}

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
