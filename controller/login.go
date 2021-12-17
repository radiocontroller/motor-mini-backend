package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"motor-mini-backend/dto"
	"motor-mini-backend/service"
	"motor-mini-backend/util"
)

func Login(ctx *gin.Context) {
	code := ctx.Query("code")
	if code == "" {
		ErrorWithMessage(ctx, 412, "code blank")
		return
	}
	login := util.MiniLogin(code)
	if login.Errcode != 0 {
		ErrorWithMessage(ctx, 400, fmt.Sprintf("loginDTO: %v", login))
		return
	}
	userDTO := dto.UserDTO{Openid: login.Openid, Unionid: login.Unionid, SessionKey: login.SessionKey}
	service.CreateOrUpdateUser(ctx, &userDTO)
	Success(ctx, map[string]string{"token": userDTO.Token})
}
