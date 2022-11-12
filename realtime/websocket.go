package realtime

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket configuration
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// Allow all origins
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
