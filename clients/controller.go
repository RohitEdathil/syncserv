package clients

import "github.com/gin-gonic/gin"

func NewController(ctx *gin.Context) {

	newsync := ClientIndexInstance.CreateNew()

	ctx.JSON(200, gin.H{
		"id":     newsync.Id,
		"secret": newsync.Secret,
	})
}

func CheckIdController(ctx *gin.Context) {

	id := ctx.Param("id")

	if ClientIndexInstance.CheckId(id) {
		ctx.JSON(200, nil)
	} else {
		ctx.JSON(404, nil)
	}
}
