package realtime

import (
	"log"
	"net/http"
	"syncserv/code"
	e "syncserv/error_handling"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// Allow all origins
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func AttachTypeSync(id string, secret string, ctx *gin.Context) {

	sharer, found := code.SyncStoreInstance.Get(id)

	if !found {
		e.PanicHTTP(e.BadRequest, "Sharer not found")
	}

	if sharer.Secret != secret {
		e.PanicHTTP(e.Unauthorized, "Invalid secret")
	}

	connection, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		e.PanicHTTP(e.BadRequest, err.Error())
	}

	go startListening(sharer, connection)

}

func startListening(sync *code.TypeSync, conn *websocket.Conn) {

	sync.Lock.Lock()
	sync.Connection = conn

	for {
		message := Message{}
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
