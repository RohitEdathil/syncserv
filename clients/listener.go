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
}

func (listener *Listener) StartListening(conn *websocket.Conn) {

	listener.Lock.Lock()
	listener.Connection = conn
	listener.Lock.Unlock()

	for {
		message := util.Message{}
		err := listener.Connection.ReadJSON(&message)

		if err != nil {
			e.PanicWS(*listener.Connection, err.Error())
			listener.Connection.Close()
			break
		}

		log.Printf("Message received from %s : %s", listener.Connection.RemoteAddr(), message)
	}

	log.Printf("Disconnected")

	listener.Of.RemoveListener(listener)

	listener.Lock.Lock()
	listener.Connection = nil
	listener.Lock.Unlock()
}
