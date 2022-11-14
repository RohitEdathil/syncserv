package handler

import (
	"syncserv/clients"
	"syncserv/codesync"
	e "syncserv/error_handling"
	"syncserv/util"
)

func HandleListenerConnected(listener *clients.Listener) {
	listener.Lock.Lock()
	codesync.SendSavedStateL(listener)
	listener.Lock.Unlock()
}

// Maps a message type to a handler
func HandleListenerMessage(listener *clients.Listener, message *util.Message) {

	listener.Lock.Lock()
	switch message.Type {

	default:
		e.PanicWS(*listener.Connection, "Invalid message type")
	}
	listener.Lock.Unlock()

}

func HandleListenerDisconnected(listener *clients.Listener) {
	listener.Lock.Lock()
	listener.Of.RemoveListener(listener)
	listener.Lock.Unlock()
}
