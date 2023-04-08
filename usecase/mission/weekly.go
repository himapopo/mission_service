package mission

import (
	"context"
	"mission_service/models"
	"mission_service/usecase/repository"
)

type WeeklyMissionUsecase interface {
	CheckMonsterKillCountMission(context.Context, *models.User) error
}

type weeklyMissionUsecase struct {
	monsterKillCountMissionRepository repository.MonsterKillCountMissionRepository
	userMissionRepository             repository.UserMissionRepository
	missionRewardUsecase              MissionRewardUsecase
}

func NewWeeklyMissionUsecase(
	monsterKillCountMissionRepository repository.MonsterKillCountMissionRepository,
	userMissionRepository repository.UserMissionRepository,
	missionRewardUsecase MissionRewardUsecase,
) weeklyMissionUsecase {
	return weeklyMissionUsecase{
		monsterKillCountMissionRepository: monsterKillCountMissionRepository,
		userMissionRepository:             userMissionRepository,
		missionRewardUsecase:              missionRewardUsecase,
	}
}

// モンスター討伐数ミッション達成チェック
func (u weeklyMissionUsecase) CheckMonsterKillCountMission(ctx context.Context, user *models.User) error {
	// ccms, err := u.monsterKillCountMissionRepository.FetchWeeklyByUserID(ctx, user.ID)
	// if err != nil {
	// 	return err
	// }
	// for _, m := range ccms {
	// 	// ミッションの達成日時更新
	// 	um := m.R.Mission.R.UserMissions[0]
	// 	um.CompletedAt = null.TimeFrom(completedAt)
	// 	if err := u.userMissionRepository.Update(ctx, um, []string{
	// 		models.UserMissionColumns.CompletedAt,
	// 		models.UserMissionColumns.UpdatedAt,
	// 	}); err != nil {
	// 		return err
	// 	}

	// 	// ミッション報酬加算
	// 	if err := u.missionRewardUsecase.ObtainRewards(ctx, um.R.User, m.R.Mission); err != nil {
	// 		return err
	// 	}

	// }
	return nil
}
