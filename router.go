package main

import (
	"syncserv/code"
	"syncserv/realtime"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.RouterGroup) {
	code.Register(router)
	realtime.Register(router)
}
