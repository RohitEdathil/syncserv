package logging

import (
	"log"
)

func Setup() {
	log.SetPrefix("[App] ")
}
