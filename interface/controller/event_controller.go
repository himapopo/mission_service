package controller

import (
	"context"
	"mission_service/usecase/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventUsecase interface {
	Login(context.Context, dto.LoginRequest) (int, error)
	MonsterKill(context.Context, dto.MonsterKillRequest) (int, error)
	MonsterLevelUp(context.Context, dto.MonsterLevelUpRequest) (int, error)
}

type eventController struct {
	eventUsecase EventUsecase
}

func NewEventController(
	eventUsecase EventUsecase,
) eventController {
	return eventController{
		eventUsecase: eventUsecase,
	}
}

func (ec eventController) Login(ctx *gin.Context) {
	var params dto.LoginRequest
	ctx.ShouldBindJSON(&params)
	status := http.StatusOK
	res := dto.Response{
		Result: true,
	}
	if s, err := ec.eventUsecase.Login(ctx, params); err != nil {
		status = s
		res.Error = err.Error()
		res.Result = false
	}
	ctx.JSON(status, res)
}

func (ec eventController) MonsterKill(ctx *gin.Context) {
	var params dto.MonsterKillRequest
	ctx.ShouldBindJSON(&params)
	status := http.StatusOK
	res := dto.Response{
		Result: true,
	}
	if s, err := ec.eventUsecase.MonsterKill(ctx, params); err != nil {
		status = s
		res.Error = err.Error()
		res.Result = false
	}
	ctx.JSON(status, res)
}

func (ec eventController) MonsterLevelUp(ctx *gin.Context) {
	var params dto.MonsterLevelUpRequest
	ctx.ShouldBindJSON(&params)
	status := http.StatusOK
	res := dto.Response{
		Result: true,
	}
	if s, err := ec.eventUsecase.MonsterLevelUp(ctx, params); err != nil {
		status = s
		res.Error = err.Error()
		res.Result = false
	}
	ctx.JSON(status, res)
}
