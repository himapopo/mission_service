package router

import "github.com/gin-gonic/gin"

type router struct {
	e               *gin.Engine
	eventController EventController
}

type EventController interface {
	Login(*gin.Context)
	MonsterKill(*gin.Context)
}

func NewRouter(
	e *gin.Engine,
	eventController EventController,
) router {
	return router{
		e:               e,
		eventController: eventController,
	}
}

func (r router) Routing() {
	r.e.POST("/login", r.eventController.Login)
	r.e.POST("/monster_kill", r.eventController.MonsterKill)
}
