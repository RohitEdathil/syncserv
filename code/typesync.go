package code

import (
	"log"
	"sync"

	e "syncserv/error_handling"
	"syncserv/util"

	"github.com/gorilla/websocket"
)

type TypeSync struct {
	Id         string
	Secret     string
	Connection *websocket.Conn
	Lock       sync.Mutex
	Listeners  []Listener
}

func (sync *TypeSync) StartListening(conn *websocket.Conn) {

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

func (sync *TypeSync) AddListener(listener Listener) {
	sync.Lock.Lock()
	sync.Listeners = append(sync.Listeners, listener)
	sync.Lock.Unlock()
}
