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
	MonsterKillCountMission(context.Context, int64, time.Time) error
}

type weeklyMissionUsecase struct {
	monsterKillCountMissionRepository repository.MonsterKillCountMissionRepository
	userMissionRepository             repository.UserMissionRepository
	userMissionProgressRepository     repository.UserMissionProgressRepository
	missionRewardUsecase              MissionRewardUsecase
	missionReleaseUsecase             MissionReleaseUsecase
}

func NewWeeklyMissionUsecase(
	monsterKillCountMissionRepository repository.MonsterKillCountMissionRepository,
	userMissionRepository repository.UserMissionRepository,
	userMissionProgressRepository repository.UserMissionProgressRepository,
	missionRewardUsecase MissionRewardUsecase,
	missionReleaseUsecase MissionReleaseUsecase,
) weeklyMissionUsecase {
	return weeklyMissionUsecase{
		monsterKillCountMissionRepository: monsterKillCountMissionRepository,
		userMissionRepository:             userMissionRepository,
		userMissionProgressRepository:     userMissionProgressRepository,
		missionRewardUsecase:              missionRewardUsecase,
		missionReleaseUsecase:             missionReleaseUsecase,
	}
}

// 任意のモンスター討伐数ミッション達成チェック
func (u weeklyMissionUsecase) MonsterKillCountMission(ctx context.Context, userID int64, requestedAt time.Time) error {
	mkcms, err := u.monsterKillCountMissionRepository.FetchWeeklyByUserID(ctx, userID)
	if err != nil {
		return err
	}

	for _, mkcm := range mkcms {
		ump := mkcm.R.Mission.R.UserMissions[0].R.UserMissionProgresses[0]
		um := mkcm.R.Mission.R.UserMissions[0]

		if ump.LastProgressUpdatedAt.Before(timeutils.WeeklyMissionResetTime(requestedAt)) {
			// 今週初めてのモンスター討伐
			ump.ProgressValue = 1
			ump.LastProgressUpdatedAt = requestedAt
			if err := u.userMissionProgressRepository.Update(ctx, ump, []string{
				models.UserMissionProgressColumns.ProgressValue,
				models.UserMissionProgressColumns.LastProgressUpdatedAt,
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
		} else if mkcm.KillCount > ump.ProgressValue {
			// 今週2体目以降のモンスター討伐
			ump.ProgressValue += 1
			ump.LastProgressUpdatedAt = requestedAt
			if err := u.userMissionProgressRepository.Update(ctx, ump, []string{
				models.UserMissionProgressColumns.ProgressValue,
				models.UserMissionProgressColumns.LastProgressUpdatedAt,
				models.UserMissionProgressColumns.UpdatedAt,
			}); err != nil {
				return err
			}
		}

		// ミッション達成済み or 達成条件に満たない場合はcontinue
		if um.CompletedAt.Valid || mkcm.KillCount > ump.ProgressValue {
			continue
		}

		// ミッションの達成日時更新
		um.CompletedAt = null.TimeFrom(requestedAt)
		if err := u.userMissionRepository.Update(ctx, um, []string{
			models.UserMissionColumns.CompletedAt,
			models.UserMissionColumns.UpdatedAt,
		}); err != nil {
			return err
		}

		mission := mkcm.R.Mission

		// ミッション解放
		if len(mission.R.CompleteMissionMissionReleases) != 0 {
			if err := u.missionReleaseUsecase.MissionRelease(ctx, userID, mission.R.CompleteMissionMissionReleases); err != nil {
				return err
			}
		}

		// ミッション報酬獲得
		if err := u.missionRewardUsecase.ObtainRewards(ctx, userID, mkcm.R.Mission, requestedAt); err != nil {
			return err
		}

	}
	return nil
}
