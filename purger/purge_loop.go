package purger

import (
	"log"
	"syncserv/clients"
	"time"
)

const PURGE_LOOP_INTERVAL = 10 * time.Second
const PURGE_TIMEOUT = 10 * time.Second

func PurgeLoop() {

	for {
		// Purge every 10 seconds
		time.Sleep(PURGE_LOOP_INTERVAL)

		log.Println("Purging")
		// Purge
		clients.ClientIndexInstance.Purge(PURGE_TIMEOUT)
	}

}
