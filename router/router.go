package router

import (
	"github.com/gin-gonic/gin"
	"github.com/growmax/noti/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/push/health", controller.Health)

	return r
}
