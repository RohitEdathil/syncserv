package mapper

import (
	"syncserv/clients"
	e "syncserv/error_handling"
	"syncserv/util"
)

// Maps a message type to a handler
func HandleListenerMessage(listener *clients.Listener, message *util.Message) {

	switch message.Type {

	default:
		e.PanicWS(*listener.Connection, "Invalid message type")
	}

}
