package codesync

import (
	"syncserv/clients"
	"syncserv/util"
)

// Syncs code of broadcaster to all listeners
func CodeSync(broadcaster *clients.Broadcaster, message *util.Message) {
	broadcaster.Lock.Lock()
	broadcaster.Text = message.Data
	// Broadcast message to all listeners
	for _, listener := range broadcaster.Listeners {
		listener.Lock.Lock()
		listener.Connection.WriteJSON(message)
		listener.Lock.Unlock()
	}

	broadcaster.Lock.Unlock()
}
