package controller

import (
	"github.com/gin-gonic/gin"
	"motor-mini-backend/core"
	"motor-mini-backend/core/logger"
	"motor-mini-backend/dto"
	"motor-mini-backend/service"
)

func Index(ctx *gin.Context) {
	markers, _ := service.ListMarker(ctx)
	Success(ctx, markers)
}

func Create(ctx *gin.Context) {
	user := dto.UserDTO{
		Openid: "测试",
		SessionKey: "fdsf",
	}
	service.CreateOrUpdateUser(ctx, &user)
	Success(ctx, user)
}

func TestRedis(ctx *gin.Context) {
	logger.Info(ctx, "TestRedis", "success")
	val, err := core.Redis.Get(ctx, "hello").Result()
	Success(ctx, map[string]interface{}{"val": val, "err": err})
}

func TestLog(ctx *gin.Context) {
	logger.Error(ctx, "成功的表现", "成功呢了的")
	service.ListMarker(ctx)
	Success(ctx, "成功")
}

//func FindUser(ctx *gin.Context) {
//	var user entity.User
//	userId, _ := strconv.Atoi(ctx.DefaultQuery("id", "0"))
//	exists, _ := service.FindUserById(ctx, uint(userId), &user)
//	logger.Warn(ctx, "warn keyworkds", user)
//	if exists <= 0 {
//		ErrorWithMessage(ctx, constant.CODE_500, "用户不存在")
//		return
//	}
//	Success(ctx, user)
//}
