package main

import (
	"mission_service/infrastructure/db"
	"mission_service/interface/controller"
	"mission_service/interface/database"
	"mission_service/router"
	"mission_service/usecase/mission"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	defer db.DB.Close()

	e := gin.Default()

	dbUtil := db.NewDB()

	// repository
	ur := database.NewuUserRepostitory(dbUtil)
	lmr := database.NewLoginMissionRepostitory(dbUtil)
	umr := database.NewUserMissionRepostitory(dbUtil)
	uir := database.NewUserItemRepostitory(dbUtil)

	// usecase
	dmuc := mission.NewDailyMissionUsecase(
		ur,
		lmr,
		umr,
		uir,
	)

	// controller
	dmc := controller.NewDailyMissionController(dmuc)

	// router
	router := router.NewRouter(e, dmc)

	router.Routing()

	e.Run()
}
