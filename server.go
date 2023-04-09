package main

import (
	"mission_service/infrastructure/db"
	"mission_service/interface/controller"
	"mission_service/interface/database"
	"mission_service/router"
	"mission_service/usecase/event"
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
	umpr := database.NewUserMissionProgressRepostitory(dbUtil)
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
		ur,
		umr,
		umpr,
		mru,
	)

	nmu := mission.NewNormailMissionUsecase(
		ccmr,
		ur,
		umr,
		umpr,
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
	
	eu := event.NewEventMissionUsecase(
		dmu,
		wmu,
		nmu,
	)

	// controller
	ec := controller.NewEventController(eu)

	// router
	router := router.NewRouter(e, ec)

	router.Routing()

	e.Run()
}
