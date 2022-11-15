package handler

import (
	"syncserv/clients"
	"syncserv/codesync"
	e "syncserv/error_handling"
	"syncserv/flag"
	"syncserv/purger"
	"syncserv/util"
)

func HandleBroadcasterConnected(broadcaster *clients.Broadcaster) {
	broadcaster.Lock.Lock()
	codesync.SendSavedStateB(broadcaster)
	flag.SendCounts(broadcaster)
	purger.ClearLastSeen(broadcaster)
	broadcaster.Lock.Unlock()
}

// Maps a message type to a handler
func HandleBroadcasterMessage(broadcaster *clients.Broadcaster, message *util.Message) {

	broadcaster.Lock.Lock()
	switch message.Type {

	case "code-state":
		codesync.CodeState(broadcaster, message)

	default:
		e.PanicWS(broadcaster.Connection, "Invalid message type")
	}
	broadcaster.Lock.Unlock()

}

func HandleBroadcasterDisconnected(broadcaster *clients.Broadcaster) {
	broadcaster.Lock.Lock()
	purger.MarkLastSeen(broadcaster)
	broadcaster.Lock.Unlock()
}
