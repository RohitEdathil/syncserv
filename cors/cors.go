package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {

	config := cors.DefaultConfig()

	// Allow all origins
	config.AllowOrigins = []string{"*"}

	return cors.New(config)
}
