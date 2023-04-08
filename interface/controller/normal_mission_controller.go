package controller

import (
	"context"
	"mission_service/usecase/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NormalMissionUsecase interface {
	MonsterKill(context.Context, dto.MonsterKillRequest) (int, error)
}

type normalMissionController struct {
	normalMissionUsecase NormalMissionUsecase
}

func NewNormalMissionController(
	normalMissionUsecase NormalMissionUsecase,
) normalMissionController {
	return normalMissionController{
		normalMissionUsecase: normalMissionUsecase,
	}
}

func (dmc normalMissionController) MonsterKill(ctx *gin.Context) {
	var params dto.MonsterKillRequest
	ctx.ShouldBindJSON(&params)
	status := http.StatusOK
	message := ""
	if s, err := dmc.normalMissionUsecase.MonsterKill(ctx, params); err != nil {
		status = s
		message = err.Error()
	}
	ctx.JSON(status, dto.Response{
		Error: message,
	})
}
