package service

import (
	"github.com/gin-gonic/gin"
	"motor-mini-backend/dao"
	"motor-mini-backend/dto"
)

func CreateOrUpdateUser(ctx *gin.Context, user *dto.UserDTO) error {
	return dao.CreateOrUpdateUser(ctx, user)
}
