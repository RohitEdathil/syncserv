package realtime

import (
	"fmt"
	"syncserv/code"
	e "syncserv/error_handling"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func AttachTypeSync(id string, secret string, ctx *gin.Context) {

	sharer, found := code.SyncStoreInstance.Get(id)

	if !found {
		e.PanicHTTP(e.InvalidRequest, "Sharer not found")
	}

	if sharer.Secret != secret {
		e.PanicHTTP(e.Unauthorized, "Invalid secret")
	}

	connection, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		e.PanicHTTP(e.InvalidRequest, "Could not upgrade connection")
	}

	sharer.Connection = connection

	fmt.Println("AttachTypeSync", id, secret)
}
