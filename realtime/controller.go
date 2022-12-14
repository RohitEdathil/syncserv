package realtime

import (
	"log"
	e "syncserv/error_handling"
	"syncserv/handler"

	"syncserv/clients"

	"github.com/gin-gonic/gin"
)

func AttachController(ctx *gin.Context) {

	// Get the id and secret from the request
	id := ctx.Param("id")
	secret := ctx.Query("secret")

	// Check if the id and secret are present
	if id == "" || secret == "" {
		e.PanicHTTP(e.BadRequest, "id and secret are required")
	}

	// Get the sharer from the index
	sharer, found := clients.ClientIndexInstance.Get(id)

	// If the sharer is not found, return 400
	if !found {
		e.PanicHTTP(e.BadRequest, "Sharer not found")
	}

	// If the secret is wrong, return 401
	if sharer.Secret != secret {
		e.PanicHTTP(e.Unauthorized, "Invalid secret")
	}

	// Upgrade the connection to a websocket
	connection, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	// Handle failure to upgrade
	if err != nil {
		log.Println(err)
		e.PanicHTTP(e.BadRequest, err.Error())
	}

	// Assign handlers
	sharer.Lock.Lock()
	sharer.ConnectedHandler = handler.HandleBroadcasterConnected
	sharer.MessageHandler = handler.HandleBroadcasterMessage
	sharer.DisconnectedHandler = handler.HandleBroadcasterDisconnected
	sharer.Lock.Unlock()

	go sharer.StartListening(connection)

}

func ListenController(ctx *gin.Context) {

	// Get id from the request
	id := ctx.Param("id")

	// Check if the id is present
	if id == "" {
		e.PanicHTTP(e.BadRequest, "id is required")
	}

	// Get the listener from the index
	sharer, found := clients.ClientIndexInstance.Get(id)

	// If the listener is not found, return 400
	if !found {
		e.PanicHTTP(e.BadRequest, "Sharer not found")
	}

	// Upgrade the connection to a websocket
	connection, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	// Handle failure to upgrade
	if err != nil {
		log.Println(err)
		e.PanicHTTP(e.BadRequest, err.Error())
	}

	// Create a new listener
	listener := clients.NewListener(sharer, connection)

	sharer.AddListener(listener)

	// Assign handler
	listener.Lock.Lock()
	listener.ConnectedHandler = handler.HandleListenerConnected
	listener.MessageHandler = handler.HandleListenerMessage
	listener.DisconnectedHandler = handler.HandleListenerDisconnected
	listener.Lock.Unlock()

	go listener.StartListening(connection)

}
