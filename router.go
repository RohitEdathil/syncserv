package main

import (
	"syncserv/clients"
	"syncserv/realtime"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.RouterGroup) {

	// Routes from each package
	clients.Register(router)
	realtime.Register(router)
}
