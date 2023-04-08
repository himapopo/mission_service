package controller

import (
	"context"
	"mission_service/usecase/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DailyMissionUsecase interface {
	Login(context.Context, dto.LoginMissionRequest) (int, error)
}

type dailyMissionController struct {
	dailyMissionUsecase DailyMissionUsecase
}

func NewDailyMissionController(
	dailyMissionUsecase DailyMissionUsecase,
) dailyMissionController {
	return dailyMissionController{
		dailyMissionUsecase: dailyMissionUsecase,
	}
}

func (dmc dailyMissionController) Login(ctx *gin.Context) {
	var params dto.LoginMissionRequest
	ctx.ShouldBindJSON(&params)
	status := http.StatusOK
	message := ""
	if s, err := dmc.dailyMissionUsecase.Login(ctx, params); err != nil {
		status = s
		message = err.Error()
	}
	ctx.JSON(status, dto.Response{
		Error: message,
	})
}
