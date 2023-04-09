package mission

import (
	"context"
	"errors"
	"fmt"
	"mission_service/models"
	"mission_service/usecase/repository"
	"mission_service/utils/timeutils"
	"time"

	"github.com/volatiletech/null/v8"
)

type DailyMissionUsecase interface {
	LoginMission(context.Context, int64, time.Time) error
}

type dailyMissionUsecase struct {
	userRepository                repository.UserRepository
	loginMissionRepository        repository.LoginMissionRepository
	userMissionRepository         repository.UserMissionRepository
	userMissionProgressRepository repository.UserMissionProgressRepository
	missionRewardUsecase          MissionRewardUsecase
	normalMissionUsecase          NormalMissionUsecase
	missionReleaseUsecase         MissionReleaseUsecase
}

func NewDailyMissionUsecase(
	userRepository repository.UserRepository,
	loginMissionRepository repository.LoginMissionRepository,
	userMissionRepository repository.UserMissionRepository,
	userMissionProgressRepository repository.UserMissionProgressRepository,
	missionRewardUsecase MissionRewardUsecase,
	normailMissionUsecase NormalMissionUsecase,
	missionReleaseUsecase MissionReleaseUsecase,
) dailyMissionUsecase {
	return dailyMissionUsecase{
		userRepository:                userRepository,
		loginMissionRepository:        loginMissionRepository,
		userMissionRepository:         userMissionRepository,
		userMissionProgressRepository: userMissionProgressRepository,
		missionRewardUsecase:          missionRewardUsecase,
		normalMissionUsecase:          normailMissionUsecase,
		missionReleaseUsecase:         missionReleaseUsecase,
	}
}

func (u dailyMissionUsecase) LoginMission(ctx context.Context, userID int64, requestedAt time.Time) error {
	lm, err := u.loginMissionRepository.FetchDailyByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if lm == nil {
		return errors.New("error: daily login mission not found")
	}

	ump := lm.R.Mission.R.UserMissions[0].R.UserMissionProgresses[0]

	// requestedAtが前回ミッション達成日時以降 &&
	// 前回ミッション達成日時がrequestedAtの日付の04:00以前の場合はミッション報酬獲得

	fmt.Println("----------------------")
	fmt.Println(ump.LastProgressUpdatedAt)
	fmt.Println(timeutils.DailyMissionResetTime(requestedAt))
	fmt.Println(ump.LastProgressUpdatedAt.Before(timeutils.DailyMissionResetTime(requestedAt)))
	if requestedAt.After(ump.LastProgressUpdatedAt) &&
		(ump.LastProgressUpdatedAt.Before(timeutils.DailyMissionResetTime(requestedAt))) {

		// ミッションの進捗更新
		ump.LastProgressUpdatedAt = requestedAt
		ump.ProgressValue = 1
		if err := u.userMissionProgressRepository.Update(ctx, ump, []string{
			models.UserMissionProgressColumns.ProgressValue,
			models.UserMissionProgressColumns.LastProgressUpdatedAt,
			models.UserMissionProgressColumns.UpdatedAt,
		}); err != nil {
			return err
		}

		// ミッション達成日時更新
		um := lm.R.Mission.R.UserMissions[0]
		um.CompletedAt = null.TimeFrom(requestedAt)
		if err := u.userMissionRepository.Update(ctx, um, []string{
			models.UserMissionColumns.CompletedAt,
			models.UserMissionColumns.UpdatedAt,
		}); err != nil {
			return err
		}

		mission := lm.R.Mission

		// ミッション報酬獲得
		if err := u.missionRewardUsecase.ObtainRewards(ctx, userID, mission); err != nil {
			return err
		}

		// ミッション解放
		if len(mission.R.CompleteMissionMissionReleases) != 0 {
			if err := u.missionReleaseUsecase.MissionRelease(ctx, userID, mission.R.CompleteMissionMissionReleases); err != nil {
				return err
			}
		}

		// userの最終ログイン日時更新 (※別システムでやる想定)
		if u.userRepository.Update(ctx, &models.User{
			ID:          userID,
			LastLoginAt: requestedAt,
		}, []string{
			models.UserColumns.LastLoginAt,
			models.UserColumns.UpdatedAt,
		}); err != nil {
			return err
		}
	}

	return nil
}
