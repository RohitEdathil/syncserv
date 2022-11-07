package code

import "github.com/gorilla/websocket"

type Listener struct {
	Of         *TypeSync
	Connection *websocket.Conn
}

func (listener *Listener) StartListening() {
	// TODO: Start here
}
