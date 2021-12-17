package service

import (
	"github.com/gin-gonic/gin"
	"motor-mini-backend/core/logger"
	"motor-mini-backend/dao"
	"motor-mini-backend/dto"
	"motor-mini-backend/entity"
)

func ListMarker(ctx *gin.Context) ([]dto.MarkerDTO, error) {
	dto, err := dao.ListMarker(ctx)
	logger.Debug(ctx, "ListMarker", dto)
	return dto, err
}

func FindMarkerById(ctx *gin.Context, id uint) (dto.MarkerDTO, error) {
	return dao.FindMarkerById(ctx, id)
}

func CreateMarker(ctx *gin.Context, marker *dto.MarkerDTO) (err error) {
	return dao.CreateMarker(ctx, marker)
}

func UpdateMarker(ctx *gin.Context, id uint, marker *entity.Marker) (err error) {
	return dao.UpdateMarker(ctx, id, marker)
}
