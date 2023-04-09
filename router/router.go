package router

import "github.com/gin-gonic/gin"

type router struct {
	e                       *gin.Engine
	dailyMissionController  DailyMissionController
	normalMissionController NormalMissionController
}

type DailyMissionController interface {
	Login(*gin.Context)
}

type NormalMissionController interface {
	MonsterKill(*gin.Context)
}

func NewRouter(
	e *gin.Engine,
	dailyMissionController DailyMissionController,
	norNormalMissionController NormalMissionController,
) router {
	return router{
		e:                       e,
		dailyMissionController:  dailyMissionController,
		normalMissionController: norNormalMissionController,
	}
}

func (r router) Routing() {
	r.e.POST("/login", r.dailyMissionController.Login)
	r.e.POST("/monster_kill", r.normalMissionController.MonsterKill)
}
