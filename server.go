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
	ccmr := database.NewCoinCountMissionRepostitory(dbUtil)
	mkmr := database.NewMonsterKillMissionRepostitory(dbUtil)
	mkkmr := database.NewMonsterKillCountMissionRepostitory(dbUtil)

	// usecase
	mru := mission.NewMissionRewardUsecase(
		ur,
		uir,
	)
	wmu := mission.NewWeeklyMissionUsecase(
		mkkmr,
		umr,
		mru,
	)
	nmu := mission.NewNormailMissionUsecase(
		ccmr,
		umr,
		mkmr,
		mru,
		wmu,
	)
	dmu := mission.NewDailyMissionUsecase(
		ur,
		lmr,
		umr,
		mru,
		nmu,
	)

	// controller
	dmc := controller.NewDailyMissionController(dmu)

	// router
	router := router.NewRouter(e, dmc)

	router.Routing()

	e.Run()
}
