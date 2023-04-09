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
	userRepository := database.NewuUserRepostitory(dbUtil)
	loginMissionRepository := database.NewLoginMissionRepostitory(dbUtil)
	userMissionRepository := database.NewUserMissionRepostitory(dbUtil)
	userMissionProgressRepository := database.NewUserMissionProgressRepostitory(dbUtil)
	userMonsterRepository := database.NewuUserMonsterRepostitory(dbUtil)
	userItemRepository := database.NewUserItemRepostitory(dbUtil)
	coinCountMissionRepository := database.NewCoinCountMissionRepostitory(dbUtil)
	monsterkillMissionRepository := database.NewMonsterKillMissionRepostitory(dbUtil)
	monsterkillCountMissionRepository := database.NewMonsterKillCountMissionRepostitory(dbUtil)
	monsterLevelUpMissionRepository := database.NewMonsterLevelUpMissionRepostitory(dbUtil)

	// usecase
	mru := mission.NewMissionRewardUsecase(
		userRepository,
		userItemRepository,
	)

	wmu := mission.NewWeeklyMissionUsecase(
		monsterkillCountMissionRepository,
		userRepository,
		userMissionRepository,
		userMissionProgressRepository,
		mru,
	)

	nmu := mission.NewNormailMissionUsecase(
		coinCountMissionRepository,
		userRepository,
		userMissionRepository,
		userMissionProgressRepository,
		userMonsterRepository,
		monsterkillMissionRepository,
		monsterLevelUpMissionRepository,
		mru,
	)

	dmu := mission.NewDailyMissionUsecase(
		userRepository,
		loginMissionRepository,
		userMissionRepository,
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
