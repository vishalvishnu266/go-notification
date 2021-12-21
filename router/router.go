package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/growmax/noti/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/push/health", controller.Health)
	r.POST("/push/getnotification", controller.GetNotification)
	r.POST("/push/addBrowser", controller.AddBrowser)
	r.GET("/push/getjwt", controller.GetJwt)
	r.GET("/push/getrecieveroption", controller.GetRecieverOption)
	r.PUT("/push/enablewebpush", controller.EnableWebpush)
	r.PUT("/push/enableemail", controller.EnableEmail)
	r.PUT("/push/enablesms", controller.EnableSms)

	return r
}
