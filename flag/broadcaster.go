package flag

import (
	"fmt"
	"syncserv/clients"
	"syncserv/util"
)

// Sends total and green flag counts
func SendCounts(broadcaster *clients.Broadcaster) {

	// Skip if broadcaster is not connected
	if broadcaster.Connection == nil {
		return
	}

	broadcaster.Connection.WriteJSON(util.Message{
		Type: "flag-count",
		Data: fmt.Sprint(len(broadcaster.Listeners)),
	})

	SendGreenFlagCount(broadcaster)
}

// Sends green flag count only
func SendGreenFlagCount(broadcaster *clients.Broadcaster) {

	broadcaster.Connection.WriteJSON(util.Message{
		Type: "green-flag-count",
		Data: fmt.Sprint(broadcaster.GreenCount),
	})
}
