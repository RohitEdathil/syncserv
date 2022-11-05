package code

import "github.com/gin-gonic/gin"

func NewController(ctx *gin.Context) {

	newsync := SyncStoreInstance.CreateNew()

	ctx.JSON(200, gin.H{
		"id":     newsync.Id,
		"secret": newsync.Secret,
	})
}
