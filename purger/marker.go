package purger

import (
	"syncserv/clients"
	"time"
)

func MarkLastSeen(broadcaster *clients.Broadcaster) {

	broadcaster.LastSeen = time.Now()

}

func ClearLastSeen(broadcaster *clients.Broadcaster) {

	broadcaster.LastSeen = time.Time{}

}
