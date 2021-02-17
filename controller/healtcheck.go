package controller

import "github.com/gin-gonic/gin"

func Healtcheck(c * gin.Context) {

	c.JSON(200, gin.H{"message": "Health is ok"})
}