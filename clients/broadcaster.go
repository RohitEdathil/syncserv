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
}

func (sync *Broadcaster) StartListening(conn *websocket.Conn) {

	sync.Lock.Lock()
	sync.Connection = conn
	sync.Lock.Unlock()

	for {
		message := util.Message{}
		err := sync.Connection.ReadJSON(&message)

		if err != nil {
			e.PanicWS(*sync.Connection, err.Error())
			sync.Connection.Close()
			break
		}

		log.Printf("Message received from %s : %s", sync.Connection.RemoteAddr(), message)
	}

	log.Printf("Disconnected")

	sync.Lock.Lock()
	sync.Connection = nil
	sync.Lock.Unlock()

}

func (sync *Broadcaster) AddListener(listener *Listener) {
	// Entry
	sync.Lock.Lock()
	listener.Lock.Lock()

	// Critical
	listener.id = sync.count
	sync.Listeners[sync.count] = *listener
	sync.count++

	// Exit
	sync.Lock.Unlock()
	listener.Lock.Unlock()
}

func (sync *Broadcaster) RemoveListener(listener *Listener) {
	sync.Lock.Lock()
	delete(sync.Listeners, listener.id)
	sync.Lock.Unlock()
}
