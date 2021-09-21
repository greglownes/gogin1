package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/greglownes/gogin1/services/pingservice"
)

type PingController interface {
	Ping(*gin.Context)
}

// holds services
type pingController struct {
	ps pingservice.PingService
}

func NewPingController(ps pingservice.PingService) PingController {
	return &pingController{ps}
}

// route handler
func (controller *pingController) Ping(c *gin.Context) {
	message, _ := controller.ps.PingMessage()
	c.JSON(200, gin.H{
		"message": message,
	})
}
