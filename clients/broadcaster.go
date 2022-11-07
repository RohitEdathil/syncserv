package clients

import (
	"log"
	"sync"

	e "syncserv/error_handling"

	"syncserv/util"

	"github.com/gorilla/websocket"
)

type Broadcaster struct {
	Id         string
	Secret     string
	Connection *websocket.Conn
	Lock       *sync.Mutex
	Listeners  map[int]Listener
	count      int

	Handler func(broadcaster *Broadcaster, message *util.Message)
}

func (broadcaster *Broadcaster) StartListening(conn *websocket.Conn) {

	broadcaster.Lock.Lock()
	broadcaster.Connection = conn
	broadcaster.Lock.Unlock()

	for {
		message := util.Message{}
		err := broadcaster.Connection.ReadJSON(&message)

		if err != nil {
			e.PanicWS(*broadcaster.Connection, err.Error())
			broadcaster.Connection.Close()
			break
		}

		log.Printf("Message received from %s : %s", broadcaster.Connection.RemoteAddr(), message)
		broadcaster.Handler(broadcaster, &message)
	}

	log.Printf("Disconnected")

	broadcaster.Lock.Lock()
	broadcaster.Connection = nil
	broadcaster.Lock.Unlock()

}

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

func (broadcaster *Broadcaster) RemoveListener(listener *Listener) {
	broadcaster.Lock.Lock()
	delete(broadcaster.Listeners, listener.id)
	broadcaster.Lock.Unlock()
}
