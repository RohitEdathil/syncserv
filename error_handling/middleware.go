package error_handling

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

// Middleware for error handling
func ErrorInterceptor(ctx *gin.Context) {
	defer Handle(ctx)
	ctx.Next()
}

// Handles the error
func Handle(ctx *gin.Context) {
	info := recover()

	// If there is no error, do nothing
	if info == nil {
		return
	}

	// If the error is from Application code, handle it
	if reflect.TypeOf(info).String() == "error_handling.HTTPError" {
		httpError := info.(HTTPError)
		ctx.AbortWithStatusJSON(httpError.Code, gin.H{
			"error": httpError.Message,
		})
		return
	}

	// If the error is from the library, don't interfere
	panic(info)

}
