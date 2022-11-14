package handler

import (
	"log"
	"syncserv/clients"
	"syncserv/codesync"
	e "syncserv/error_handling"
	"syncserv/util"
)

func HandleBroadcasterConnected(broadcaster *clients.Broadcaster) {
	broadcaster.Lock.Lock()
	codesync.SendSavedStateB(broadcaster)
	broadcaster.Lock.Unlock()
}

// Maps a message type to a handler
func HandleBroadcasterMessage(broadcaster *clients.Broadcaster, message *util.Message) {

	broadcaster.Lock.Lock()
	switch message.Type {

	case "code-state":
		codesync.CodeState(broadcaster, message)

	default:
		e.PanicWS(*broadcaster.Connection, "Invalid message type")
	}
	broadcaster.Lock.Unlock()

}

func HandleBroadcasterDisconnected(broadcaster *clients.Broadcaster) {
	broadcaster.Lock.Lock()
	// Remove broadcaster from list if no more listeners
	if len(broadcaster.Listeners) == 0 {
		log.Printf("No more listeners, closing broadcaster")
		clients.ClientIndexInstance.Delete(broadcaster.Id)
	}
	broadcaster.Lock.Unlock()
}
