package error_handling

type HTTPError struct {
	Code    int
	Message string
}

func PanicHTTP(code int, message string) {
	panic(HTTPError{
		Code:    code,
		Message: message,
	})
}
