package clients

import (
	"log"
	"sync"
	e "syncserv/error_handling"
	"syncserv/util"

	"github.com/gorilla/websocket"
)

// Listens for messages from the broadcaster, and sends messages to the broadcaster
type Listener struct {
	id         int
	Of         *Broadcaster
	Connection *websocket.Conn
	Lock       *sync.Mutex
	GreenFlag  bool

	ConnectedHandler    func(listener *Listener)
	MessageHandler      func(listener *Listener, message *util.Message)
	DisconnectedHandler func(listener *Listener)
}

// Creates a new listener
func NewListener(of *Broadcaster, connection *websocket.Conn) *Listener {
	return &Listener{
		Of:         of,
		Connection: connection,
		Lock:       &sync.Mutex{},
		GreenFlag:  true,
	}
}

// Starts listening for messages from the listener
func (listener *Listener) StartListening(conn *websocket.Conn) {

	// Assign connection
	listener.Lock.Lock()
	listener.Connection = conn
	listener.Lock.Unlock()

	listener.ConnectedHandler(listener)

	// Listen loop
	for {
		// Readig and parsing message
		message := util.Message{}
		err := listener.Connection.ReadJSON(&message)

		// Error handling
		if err != nil {
			listener.Lock.Lock()
			e.PanicWS(*listener.Connection, err.Error())
			listener.Lock.Unlock()
			listener.Connection.Close()
			break
		}

		// Handling message
		log.Printf("Message received from %s", listener.Connection.RemoteAddr())
		listener.MessageHandler(listener, &message)
	}

	// Broadcast disconnection
	log.Printf("Disconnected")

	listener.DisconnectedHandler(listener)

	listener.Lock.Lock()
	listener.Connection = nil
	listener.Lock.Unlock()
}
