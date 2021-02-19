package controller

import (
	"aws-test/schema"
	"aws-test/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendMsg(c *gin.Context) {
	var response schema.Response

	if err := c.ShouldBindJSON(&response); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	services.SendSNSMsg(response)
}
