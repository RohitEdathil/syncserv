package codesync

import (
	"syncserv/clients"
	"syncserv/util"
)

// Syncs code of broadcaster to all listeners
func CodeState(broadcaster *clients.Broadcaster, message *util.Message) {
	broadcaster.Text = message.Data
	// Broadcast message to all listeners
	for _, listener := range broadcaster.Listeners {
		listener.Lock.Lock()
		listener.Connection.WriteJSON(message)
		listener.Lock.Unlock()
	}

}

func SendSavedStateB(broadcaster *clients.Broadcaster) {
	broadcaster.Connection.WriteJSON(util.Message{
		Type: "code-state",
		Data: broadcaster.Text,
	})
}

func SendSavedStateL(listener *clients.Listener) {
	listener.Connection.WriteJSON(util.Message{
		Type: "code-state",
		Data: listener.Of.Text,
	})
}
