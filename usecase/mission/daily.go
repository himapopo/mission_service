package mission

import (
	"context"
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
	userRepository         repository.UserRepository
	loginMissionRepository repository.LoginMissionRepository
	userMissionRepository  repository.UserMissionRepository
	missionRewardUsecase   MissionRewardUsecase
	normalMissionUsecase   NormalMissionUsecase
}

func NewDailyMissionUsecase(
	userRepository repository.UserRepository,
	loginMissionRepository repository.LoginMissionRepository,
	userMissionRepository repository.UserMissionRepository,
	missionRewardUsecase MissionRewardUsecase,
	normailMissionUsecase NormalMissionUsecase,
) dailyMissionUsecase {
	return dailyMissionUsecase{
		userRepository:         userRepository,
		loginMissionRepository: loginMissionRepository,
		userMissionRepository:  userMissionRepository,
		missionRewardUsecase:   missionRewardUsecase,
		normalMissionUsecase:   normailMissionUsecase,
	}
}

func (u dailyMissionUsecase) LoginMission(ctx context.Context, userID int64, requestedAt time.Time) error {
	lm, err := u.loginMissionRepository.FetchByUserIDAndLoginCount(ctx, userID, 1)
	if err != nil {
		return err
	}
	// 前回ミッション達成日時が今日の04:00以前の場合はミッション報酬獲得

	if lm != nil &&
		len(lm.R.Mission.R.UserMissions) != 0 &&
		(lm.R.Mission.R.UserMissions[0].CompletedAt.Time.Before(timeutils.DailyMissionResetTime(requestedAt))) {

		// ミッション達成日時更新
		um := lm.R.Mission.R.UserMissions[0]
		um.CompletedAt = null.TimeFrom(requestedAt)
		if err := u.userMissionRepository.Update(ctx, um, []string{
			models.UserMissionColumns.CompletedAt,
			models.UserMissionColumns.UpdatedAt,
		}); err != nil {
			return err
		}

		// ミッション報酬獲得
		if err := u.missionRewardUsecase.ObtainRewards(ctx, userID, lm.R.Mission); err != nil {
			return err
		}

		// userの最終ログイン日時更新 (別システムでやる想定でいいかも)
		// if u.userRepository.Update(ctx, &models.User{
		// 	ID:          params.UserID,
		// 	LastLoginAt: params.RequestedAt,
		// }, []string{
		// 	models.UserColumns.LastLoginAt,
		// 	models.UserColumns.UpdatedAt,
		// }); err != nil {
		// 	return err
		// }
	}

	return nil
}
