package realtime

import (
	"fmt"
	"syncserv/code"
	e "syncserv/error_handling"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func AttachTypeSync(id string, secret string) {

	sharer, found := code.SyncStoreInstance.Get(id)

	if !found {
		e.PanicHTTP(e.InvalidRequest, "Sharer not found")
	}

	if sharer.Secret != secret {
		e.PanicHTTP(e.Unauthorized, "Invalid secret")
	}

	fmt.Println("AttachTypeSync", id, secret)
}
