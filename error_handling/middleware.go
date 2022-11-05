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

func Handle(ctx *gin.Context) {
	info := recover()

	if info == nil {
		return
	}

	if reflect.TypeOf(info).String() == "error_handling.HTTPError" {
		httpError := info.(HTTPError)
		ctx.AbortWithStatusJSON(httpError.Code, gin.H{
			"error": httpError.Message,
		})
		return
	}

	panic(info)

}
