package mission

import (
	"context"
	"mission_service/infrastructure/db"
	"mission_service/models"
	"mission_service/usecase/dto"
	"mission_service/usecase/repository"
	"mission_service/utils/timeutils"
	"net/http"

	"github.com/volatiletech/null/v8"
)

type dailyMissionUsecase struct {
	userRepository         repository.UserRepository
	loginMissionRepository repository.LoginMissionRepository
	userMissionRepository  repository.UserMissionRepository
	missionRewardUsecase   MissionRewardUsecase
	normailMissionUsecase  NormailMissionUsecase
}

func NewDailyMissionUsecase(
	userRepository repository.UserRepository,
	loginMissionRepository repository.LoginMissionRepository,
	userMissionRepository repository.UserMissionRepository,
	missionRewardUsecase MissionRewardUsecase,
	normailMissionUsecase NormailMissionUsecase,
) dailyMissionUsecase {
	return dailyMissionUsecase{
		userRepository:         userRepository,
		loginMissionRepository: loginMissionRepository,
		userMissionRepository:  userMissionRepository,
		missionRewardUsecase:   missionRewardUsecase,
		normailMissionUsecase:  normailMissionUsecase,
	}
}

func (u dailyMissionUsecase) Login(ctx context.Context, params dto.LoginMissionRequest) (int, error) {
	lm, err := u.loginMissionRepository.FetchByUserIDAndLoginCount(ctx, params.UserID, 1)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	// 前回ミッション達成日時が今日の04:00以前の場合はミッション報酬獲得
	if err := db.InTx(ctx, func(ctx context.Context) error {

		if lm != nil &&
			len(lm.R.Mission.R.UserMissions) != 0 &&
			(lm.R.Mission.R.UserMissions[0].CompletedAt.Time.Before(timeutils.DailyMissionResetTime())) {

			// ログインミッション達成日時更新
			um := lm.R.Mission.R.UserMissions[0]
			um.CompletedAt = null.TimeFrom(params.RequestedAt)
			if err := u.userMissionRepository.Update(ctx, um, []string{
				models.UserMissionColumns.CompletedAt,
				models.UserMissionColumns.UpdatedAt,
			}); err != nil {
				return err
			}

			user := um.R.User
			// ログインミッション報酬獲得
			if err := u.missionRewardUsecase.ObtainRewards(ctx, user, lm.R.Mission); err != nil {
				return err
			}

			// コイン獲得枚数ミッション達成チェック
			if err := u.normailMissionUsecase.CheckCoinCountMission(ctx, user, params.RequestedAt); err != nil {
				return err
			}

			// userの最終ログイン日時更新 (別システムでやる想定でいいかも)
			if u.userRepository.Update(ctx, &models.User{
				ID:          params.UserID,
				LastLoginAt: params.RequestedAt,
			}, []string{
				models.UserColumns.LastLoginAt,
				models.UserColumns.UpdatedAt,
			}); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
