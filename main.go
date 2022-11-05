package main

import (
	"syncserv/error_handling"

	"github.com/gin-gonic/gin"
)

func main() {
	serv := gin.Default()

	serv.Use(error_handling.ErrorInterceptor)

	serv.SetTrustedProxies(nil)

	registerRoutes(serv.Group("/api"))

	serv.Run(":8080")
}
