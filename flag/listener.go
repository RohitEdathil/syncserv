package flag

import "syncserv/clients"

// Updates the green flag count for a broadcaster and sends it to broadcaster
func ListenerConnected(listener *clients.Listener) {
	listener.Of.Lock.Lock()
	listener.Of.GreenCount++
	SendCounts(listener.Of)
	listener.Of.Lock.Unlock()
}

// Switches flag for a listener and sends green flag count to broadcaster
func ListenerFlagSwitched(listener *clients.Listener) {
	listener.Of.Lock.Lock()

	// Switch flag
	listener.GreenFlag = !listener.GreenFlag

	// Update green flag count
	if listener.GreenFlag {
		listener.Of.GreenCount++
	} else {
		listener.Of.GreenCount--
	}
	SendGreenFlagCount(listener.Of)
	listener.Of.Lock.Unlock()
}

// Updates the green flag count for a broadcaster and sends it to broadcaster
func ListenerDisconnected(listener *clients.Listener) {
	listener.Of.Lock.Lock()

	// Update green flag count only if listener had green flag
	if listener.GreenFlag {
		listener.Of.GreenCount--
	}

	SendCounts(listener.Of)
	listener.Of.Lock.Unlock()
}
