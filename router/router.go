package router

import (
	"github.com/gin-gonic/gin"
	"github.com/growmax/noti/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/push/health", controller.Health)
	r.GET("/push/createuser", controller.CreateUser)
	r.GET("/push/getusers", controller.GetAllUser)
	r.POST("/push/addBrowser", controller.AddBrowser)

	return r
}
