package server

import (
	"aws-test/controller"
	"github.com/gin-gonic/gin"
)

func RouterInit(engine *gin.Engine) {
	group := engine.Group("/aws")

	group.GET("/health-check", controller.Healtcheck)
	group.POST("/send-msg", controller.SendMsg)
}
