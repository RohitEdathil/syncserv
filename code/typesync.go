package code

import "github.com/gorilla/websocket"

type TypeSync struct {
	Id         string
	Secret     string
	Connection *websocket.Conn
}
