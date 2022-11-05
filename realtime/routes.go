package realtime

import "github.com/gin-gonic/gin"

func Register(router *gin.RouterGroup) {
	router = router.Group("/ws")

	router.GET("attach", AttachController)

}
