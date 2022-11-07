package realtime

import (
	"log"
	"syncserv/code"
	e "syncserv/error_handling"

	"github.com/gin-gonic/gin"
)

func AttachController(ctx *gin.Context) {

	id := ctx.Param("id")
	secret := ctx.Query("secret")

	if id == "" || secret == "" {
		e.PanicHTTP(e.BadRequest, "id and secret are required")
	}

	sharer, found := code.SyncStoreInstance.Get(id)

	if !found {
		e.PanicHTTP(e.BadRequest, "Sharer not found")
	}

	if sharer.Secret != secret {
		e.PanicHTTP(e.Unauthorized, "Invalid secret")
	}

	connection, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		log.Println(err)
		e.PanicHTTP(e.BadRequest, err.Error())
	}

	go sharer.StartListening(connection)

}

func ListenController(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		e.PanicHTTP(e.BadRequest, "id is required")
	}

	sharer, found := code.SyncStoreInstance.Get(id)

	if !found {
		e.PanicHTTP(e.BadRequest, "Sharer not found")
	}

	connection, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		log.Println(err)
		e.PanicHTTP(e.BadRequest, err.Error())
	}

	listener := code.Listener{
		Of:         sharer,
		Connection: connection,
	}

	sharer.AddListener(listener)

	go listener.StartListening()

}
