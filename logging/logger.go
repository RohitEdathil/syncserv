package logging

import (
	"log"
)

// Sets up the logger
func Setup() {

	// Prefix App logs with [APP]
	log.SetPrefix("[App] ")
}
