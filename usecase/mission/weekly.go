package mission

import (
	"context"
	"mission_service/models"
	"mission_service/usecase/repository"
	"mission_service/utils/timeutils"
	"time"

	"github.com/volatiletech/null/v8"
)

type WeeklyMissionUsecase interface {
	CheckMonsterKillCountMission(context.Context, int64, time.Time) error
}

type weeklyMissionUsecase struct {
	monsterKillCountMissionRepository repository.MonsterKillCountMissionRepository
	userRepository                    repository.UserRepository
	userMissionRepository             repository.UserMissionRepository
	userMissionProgressRepository     repository.UserMissionProgressRepository
	missionRewardUsecase              MissionRewardUsecase
}

func NewWeeklyMissionUsecase(
	monsterKillCountMissionRepository repository.MonsterKillCountMissionRepository,
	userRepository repository.UserRepository,
	userMissionRepository repository.UserMissionRepository,
	userMissionProgressRepository repository.UserMissionProgressRepository,
	missionRewardUsecase MissionRewardUsecase,
) weeklyMissionUsecase {
	return weeklyMissionUsecase{
		monsterKillCountMissionRepository: monsterKillCountMissionRepository,
		userRepository:                    userRepository,
		userMissionRepository:             userMissionRepository,
		userMissionProgressRepository:     userMissionProgressRepository,
		missionRewardUsecase:              missionRewardUsecase,
	}
}

// 任意のモンスター討伐数ミッション達成チェック
func (u weeklyMissionUsecase) CheckMonsterKillCountMission(ctx context.Context, userID int64, completedAt time.Time) error {
	mkcms, err := u.monsterKillCountMissionRepository.FetchWeeklyByUserID(ctx, userID)
	if err != nil {
		return err
	}

	for _, mkcm := range mkcms {
		ump := mkcm.R.Mission.R.UserMissions[0].R.UserMissionProgresses[0]
		um := mkcm.R.Mission.R.UserMissions[0]
		if ump.UpdatedAt.Before(timeutils.WeeklyMissionResetTime(completedAt)) {
			// 今週初めてのモンスター討伐
			ump.ProgressValue = 1
			if err := u.userMissionProgressRepository.Update(ctx, ump, []string{
				models.UserMissionProgressColumns.ProgressValue,
				models.UserMissionProgressColumns.UpdatedAt,
			}); err != nil {
				return err
			}

			um.CompletedAt = null.TimeFromPtr(nil)
			if err := u.userMissionRepository.Update(ctx, um, []string{
				models.UserMissionColumns.CompletedAt,
				models.UserMissionColumns.UpdatedAt,
			}); err != nil {
				return err
			}
		} else {
			// 今週2体目以降のモンスター討伐
			ump.ProgressValue += 1
			if err := u.userMissionProgressRepository.Update(ctx, ump, []string{
				models.UserMissionProgressColumns.ProgressValue,
				models.UserMissionProgressColumns.UpdatedAt,
			}); err != nil {
				return err
			}
		}

		// ミッションの達成条件に満たない場合
		if mkcm.KillCount > ump.ProgressValue {
			continue
		}

		// ミッションの達成日時更新
		um.CompletedAt = null.TimeFrom(completedAt)
		if err := u.userMissionRepository.Update(ctx, um, []string{
			models.UserMissionColumns.CompletedAt,
			models.UserMissionColumns.UpdatedAt,
		}); err != nil {
			return err
		}

		// ミッション報酬獲得
		if err := u.missionRewardUsecase.ObtainRewards(ctx, userID, mkcm.R.Mission); err != nil {
			return err
		}

	}
	return nil
}
