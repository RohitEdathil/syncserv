package error_handling

import (
	"github.com/gorilla/websocket"
)

type HTTPError struct {
	Code    int
	Message string
}

type WSError struct {
	Message string `json:"error"`
}

func PanicHTTP(code int, message string) {
	panic(HTTPError{
		Code:    code,
		Message: message,
	})
}

func PanicWS(ws websocket.Conn, message string) {
	ws.WriteJSON(WSError{message})
}
