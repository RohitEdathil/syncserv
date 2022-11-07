package realtime

import (
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
		e.PanicHTTP(e.BadRequest, err.Error())
	}

	go sharer.StartListening(connection)

}
