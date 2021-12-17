package dao

import (
	"errors"
	"motor-mini-backend/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrUpdateUser(ctx *gin.Context, user *dto.UserDTO) error {
	result := DB(ctx).Table("users").Where("openid = ?", user.Openid).Take(user)
	if result.Error == nil {
		return DB(ctx).Table("users").Where("openid = ?", user.Openid).Updates(user).Error
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return DB(ctx).Table("users").Create(user).Error
	}
	return result.Error
}
