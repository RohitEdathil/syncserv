package clients

import "github.com/gin-gonic/gin"

func Register(router *gin.RouterGroup) {

	router = router.Group("/code")

	router.GET("new", NewController)
	router.GET(":id", CheckIdController)
}
