package main

import (
	"syncserv/cors"
	"syncserv/error_handling"
	"syncserv/logging"

	"github.com/gin-gonic/gin"
)

func main() {
	serv := gin.Default()

	logging.Setup()

	serv.Use(error_handling.ErrorInterceptor)
	serv.Use(cors.CorsMiddleware())

	serv.SetTrustedProxies(nil)

	registerRoutes(serv.Group("/api"))

	serv.Run(":8080")
}
