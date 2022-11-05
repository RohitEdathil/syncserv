package realtime

import (
	"syncserv/error_handling"

	"github.com/gin-gonic/gin"
)

func AttachController(ctx *gin.Context) {

	id := ctx.Query("id")
	secret := ctx.Query("secret")

	if id == "" || secret == "" {
		error_handling.PanicHTTP(error_handling.InvalidRequest, "id and secret are required")
	}

	AttachTypeSync(id, secret, ctx)

}
