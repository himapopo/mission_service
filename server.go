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
	userRepository := database.NewUserRepository(dbUtil)
	loginMissionRepository := database.NewLoginMissionRepository(dbUtil)
	userMissionRepository := database.NewUserMissionRepository(dbUtil)
	userMissionProgressRepository := database.NewUserMissionProgressRepository(dbUtil)
	userMonsterRepository := database.NewuUserMonsterRepository(dbUtil)
	userItemRepository := database.NewUserItemRepository(dbUtil)
	coinCountMissionRepository := database.NewCoinCountMissionRepository(dbUtil)
	monsterkillMissionRepository := database.NewMonsterKillMissionRepository(dbUtil)
	monsterkillCountMissionRepository := database.NewMonsterKillCountMissionRepository(dbUtil)
	monsterLevelUpMissionRepository := database.NewMonsterLevelUpMissionRepository(dbUtil)
	monsterLevelUpCountMissionRepository := database.NewMonsterLevelUpCountMissionRepository(dbUtil)
	getItemMissionRepository := database.NewGetItemMissionRepository(dbUtil)

	// usecase
	missionReleaseUsecase := mission.NewMissionReleaseUsecase(
		userMissionRepository,
		userMissionProgressRepository,
	)

	missionRewardUsecase := mission.NewMissionRewardUsecase(
		userRepository,
		userItemRepository,
	)

	weeklyMissionUsecase := mission.NewWeeklyMissionUsecase(
		monsterkillCountMissionRepository,
		userRepository,
		userMissionRepository,
		userMissionProgressRepository,
		missionRewardUsecase,
		missionReleaseUsecase,
	)

	normalMissionUsecase := mission.NewNormailMissionUsecase(
		coinCountMissionRepository,
		userRepository,
		userItemRepository,
		userMissionRepository,
		userMissionProgressRepository,
		userMonsterRepository,
		monsterkillMissionRepository,
		monsterLevelUpMissionRepository,
		monsterLevelUpCountMissionRepository,
		getItemMissionRepository,
		missionRewardUsecase,
		missionReleaseUsecase,
	)

	dailyMissionUsecase := mission.NewDailyMissionUsecase(
		userRepository,
		loginMissionRepository,
		userMissionRepository,
		userMissionProgressRepository,
		missionRewardUsecase,
		normalMissionUsecase,
		missionReleaseUsecase,
	)

	eu := event.NewEventMissionUsecase(
		dailyMissionUsecase,
		weeklyMissionUsecase,
		normalMissionUsecase,
	)

	// controller
	ec := controller.NewEventController(eu)

	// router
	router := router.NewRouter(e, ec)

	router.Routing()

	e.Run()
}
