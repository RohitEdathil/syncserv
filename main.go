package main

import (
	"syncserv/cors"
	"syncserv/error_handling"
	"syncserv/logging"
	"syncserv/purger"

	"github.com/gin-gonic/gin"
)

func main() {
	serv := gin.Default()

	// Initializing stuff
	logging.Setup()
	serv.SetTrustedProxies(nil)

	// Middleware
	serv.Use(error_handling.ErrorInterceptor)
	serv.Use(cors.CorsMiddleware())

	// Run purger concurrently
	go purger.PurgeLoop()

	// Routing entry point
	registerRoutes(serv.Group("/api"))

	serv.Run(":8080")
}
