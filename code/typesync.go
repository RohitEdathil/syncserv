package code

import (
	"sync"

	"github.com/gorilla/websocket"
)

type TypeSync struct {
	Id         string
	Secret     string
	Connection *websocket.Conn
	Lock       sync.Mutex
}
