package clients

import (
	"log"
	"sync"
	e "syncserv/error_handling"
	"syncserv/util"

	"github.com/gorilla/websocket"
)

type Listener struct {
	id         int
	Of         *Broadcaster
	Connection *websocket.Conn
	Lock       *sync.Mutex

	Handler func(listener *Listener, message *util.Message)
}

func (listener *Listener) StartListening(conn *websocket.Conn) {

	listener.Lock.Lock()
	listener.Connection = conn
	listener.Lock.Unlock()

	for {
		message := util.Message{}
		err := listener.Connection.ReadJSON(&message)

		if err != nil {
			listener.Lock.Lock()
			e.PanicWS(*listener.Connection, err.Error())
			listener.Lock.Unlock()
			listener.Connection.Close()
			break
		}

		log.Printf("Message received from %s : %s", listener.Connection.RemoteAddr(), message)
		listener.Handler(listener, &message)
	}

	log.Printf("Disconnected")

	listener.Of.RemoveListener(listener)

	listener.Lock.Lock()
	listener.Connection = nil
	listener.Lock.Unlock()
}
