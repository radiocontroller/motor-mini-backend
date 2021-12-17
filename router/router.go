package router

import (
	"motor-mini-backend/controller"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(router *gin.Engine) {

	// router.NoRoute(controller.NOTFOUND)
	// router.GET("/", controller.Index)
	//router.GET("/create", controller.Create)
	//router.GET("/test", controller.TestRedis)
	// router.GET("/testlog", controller.TestLog)
	//router.GET("/find", controller.FindUser)

	api := router.Group("/api")
	{
		api.GET("/login", controller.Login)
		api.GET("/markers", controller.MarkerIndex)
		//api.GET("/markers/:id", controller.MarkerShow)
		//api.POST("/markers", controller.MarkerCreate)
		//api.PUT("/markers/:id", controller.MarkerUpdate)
	}
}
