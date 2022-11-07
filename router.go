package main

import (
	"syncserv/clients"
	"syncserv/realtime"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.RouterGroup) {
	clients.Register(router)
	realtime.Register(router)
}
