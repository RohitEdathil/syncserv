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
	Channel    <-chan interface{}
	Listeners  []Listener
}

func (sync *TypeSync) StartListening(conn *websocket.Conn) {

	sync.Lock.Lock()
	sync.Connection = conn

	for {
		message := util.Message{}
		err := sync.Connection.ReadJSON(&message)

		if err != nil {
			e.PanicWS(*sync.Connection, err.Error())
			sync.Connection.Close()
			break
		}

		log.Println("Message received", message)
	}

	log.Printf("Disconnected")
	sync.Connection = nil
	sync.Lock.Unlock()

}
