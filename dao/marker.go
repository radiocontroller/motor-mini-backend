package dao

import (
	"motor-mini-backend/dto"
	"motor-mini-backend/entity"
	"github.com/gin-gonic/gin"
)

func ListMarker(ctx *gin.Context) ([]dto.MarkerDTO, error) {
	var markers []dto.MarkerDTO

	result := DB(ctx).Table("markers").Select("*").Find(&markers)
	return markers, result.Error
}

func FindMarkerById(ctx *gin.Context, id uint) (dto.MarkerDTO, error) {
	var marker dto.MarkerDTO
	result := DB(ctx).Table("markers").First(&marker, id)
	return marker, result.Error
}

func CreateMarker(ctx *gin.Context, marker *dto.MarkerDTO) error {
	return DB(ctx).Table("markers").Create(marker).Error
}

func UpdateMarker(ctx *gin.Context, id uint, marker *entity.Marker) error {
	return DB(ctx).Table("markers").Where("id = ?", id).Updates(*marker).Error
}
