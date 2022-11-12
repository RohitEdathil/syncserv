package mapper

import (
	"syncserv/clients"
	"syncserv/codesync"
	e "syncserv/error_handling"
	"syncserv/util"
)

// Maps a message type to a handler
func HandleBroadcasterMessage(broadcaster *clients.Broadcaster, message *util.Message) {

	switch message.Type {

	case "code-diff":
		codesync.CodeSync(broadcaster, message)

	default:
		e.PanicWS(*broadcaster.Connection, "Invalid message type")
	}

}
