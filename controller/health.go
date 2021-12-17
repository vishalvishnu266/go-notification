package controller

import (
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hey I am alive :-)",
	})
}
