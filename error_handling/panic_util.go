package error_handling

import (
	"github.com/gorilla/websocket"
)

// HTTPError is a wrapper for HTTP errors
type HTTPError struct {
	Code    int
	Message string
}

// WSError is a wrapper for WebSocket errors
type WSError struct {
	Message string `json:"error"`
}

// Utility function for panicking with HTTPError
func PanicHTTP(code int, message string) {
	panic(HTTPError{
		Code:    code,
		Message: message,
	})
}

// Utility function for panicking with WebSocketError
func PanicWS(ws websocket.Conn, message string) {
	ws.WriteJSON(WSError{message})
}
