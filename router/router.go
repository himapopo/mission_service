package router

import "github.com/gin-gonic/gin"

type router struct {
	e                      *gin.Engine
	dailyMissionController DailyMissionController
}

type DailyMissionController interface {
	Login(*gin.Context)
}

func NewRouter(
	e *gin.Engine,
	DailyMissionController DailyMissionController,
) router {
	return router{
		e:                      e,
		dailyMissionController: DailyMissionController,
	}
}

func (r router) Routing() {
	r.e.POST("/login", r.dailyMissionController.Login)
}
