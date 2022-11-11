package codesync

import (
	"syncserv/clients"
	"syncserv/util"
)

func CodeSync(broadcaster *clients.Broadcaster, message *util.Message) {
	broadcaster.Lock.Lock()

	for _, listener := range broadcaster.Listeners {
		listener.Lock.Lock()
		listener.Connection.WriteJSON(message)
		listener.Lock.Unlock()
	}

	broadcaster.Lock.Unlock()
}
