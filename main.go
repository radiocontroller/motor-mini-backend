package main

import (
	"motor-mini-backend/config"
	"motor-mini-backend/core"
	"motor-mini-backend/core/logger"
	"motor-mini-backend/middleware"
	"motor-mini-backend/router"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfig() //加载配置
}

func main() {

	logger.InitLog() //日志初始化
	logger.SyncLog()

	core.ConnectDB()    //连接数据库
	core.ConnectRedis() //连接redis

	if config.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	middleware.LoadMiddleware(r) //加载中间件

	router.LoadRoutes(r) //加载路由

	r.Run(config.App.Port) //启动服务

	resourceRelease()
}

func resourceRelease() {
	go func() {
		core.DisconnectDB()
	}()
}
