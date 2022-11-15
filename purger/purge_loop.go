package purger

import (
	"log"
	"syncserv/clients"
	"time"
)

func PurgeLoop() {

	for {
		// Purge every 10 seconds
		time.Sleep(PURGE_LOOP_INTERVAL)

		log.Println("Purging")
		// Purge
		clients.ClientIndexInstance.Purge(PURGE_TIMEOUT)
	}

}
