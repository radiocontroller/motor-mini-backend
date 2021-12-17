package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"motor-mini-backend/dto"
	"motor-mini-backend/entity"
	"motor-mini-backend/service"
	"strconv"
)

type TokenHeader struct {
	Authorization   string    `header:"Authorization" binding:"required"`
}

func MarkerIndex(ctx *gin.Context) {
	h := TokenHeader{}
	if err := ctx.ShouldBindHeader(&h); err != nil {
		ErrorWithMessage(ctx, 412, "token blank")
		return
	}
	fmt.Println("header:", h)
	if users, err := service.ListMarker(ctx); err != nil {
		ErrorWithMessage(ctx, 500, fmt.Sprintf("failed: %v", err))
		return
	} else {
		Success(ctx, users)
	}
}

func MarkerShow(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ErrorWithMessage(ctx, 412, err.Error())
	} else {
		marker, _ := service.FindMarkerById(ctx, uint(id))
		Success(ctx, marker)
	}
}

func MarkerCreate(ctx *gin.Context) {
	var form dto.MarkerDTO
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ErrorWithMessage(ctx, 412, err.Error())
		return
	}
	if err := service.CreateMarker(ctx, &form); err != nil {
		ErrorWithMessage(ctx, 500, fmt.Sprintf("create failed: %v", err))
	} else {
		Success(ctx, []int{})
	}
}

func MarkerUpdate(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var form entity.Marker
	if err := ctx.Bind(&form); err != nil {
		ErrorWithMessage(ctx, 500, fmt.Sprintf("update ShouldBindJSON fail: %v", err))
		return
	}
	if err := service.UpdateMarker(ctx, uint(id), &form); err != nil {
		ErrorWithMessage(ctx, 500, fmt.Sprintf("update fail: %v", err))
	} else {
		Success(ctx, []int{})
	}
}
