package dao

import (
	core "motor-mini-backend/core"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DB(ctx *gin.Context) *gorm.DB {
	return core.DB(ctx)
}
