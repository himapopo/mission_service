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
	userItemRepository     repository.UserItemRepository
}

func NewDailyMissionUsecase(
	userRepository repository.UserRepository,
	loginMissionRepository repository.LoginMissionRepository,
	userMissionRepository repository.UserMissionRepository,
	userItemRepository repository.UserItemRepository,
) dailyMissionUsecase {
	return dailyMissionUsecase{
		userRepository:         userRepository,
		loginMissionRepository: loginMissionRepository,
		userMissionRepository:  userMissionRepository,
		userItemRepository:     userItemRepository,
	}
}

func (u dailyMissionUsecase) Login(ctx context.Context, params dto.LoginMissionRequest) (int, error) {
	lm, err := u.loginMissionRepository.FetchByUserIDAndLoginCount(ctx, params.UserID, 1)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	// 前回ミッション達成日時が今日の04:00以前の場合はミッション報酬獲得
	if lm != nil &&
		len(lm.R.Mission.R.UserMissions) != 0 &&
		(lm.R.Mission.R.UserMissions[0].CompletedAt.Time.Before(timeutils.DailyMissionResetTime())) {
		if _, err := db.InTx(ctx, func(c context.Context) (interface{}, error) {

			// ログインミッションの達成日時更新
			um := lm.R.Mission.R.UserMissions[0]
			um.CompletedAt = null.TimeFrom(params.RequestedAt)
			if err := u.userMissionRepository.Update(ctx, um, []string{
				models.UserMissionColumns.CompletedAt,
				models.UserMissionColumns.UpdatedAt,
			}); err != nil {
				return nil, err
			}

			// TODO: 報酬獲得は別のusecaseに切り出す。mission報酬加算
			user := um.R.User
			mission := lm.R.Mission
			rewardCoins := mission.R.MissionRewardCoins
			rewardItems := mission.R.MissionRewardItems

			// コイン報酬加算
			if len(rewardCoins) != 0 {
				user.CoinCount += rewardCoins[0].CoinCount
				if u.userRepository.Update(ctx, user, []string{
					models.UserColumns.CoinCount,
					models.UserColumns.UpdatedAt,
				}); err != nil {
					return http.StatusInternalServerError, err
				}
			}

			// アイテム報酬加算
			for _, item := range rewardItems {
				ui, err := u.userItemRepository.FetchByItemIDAndUserID(ctx, params.UserID, item.ItemID)
				if err != nil {
					return http.StatusInternalServerError, err
				}
				if ui == nil {
					// 所持していないアイテムの場合はレコード作成
					if err := u.userItemRepository.Create(ctx, &models.UserItem{
						UserID: params.UserID,
						ItemID: item.ItemID,
						Count:  int(item.ItemCount),
					}); err != nil {
						return http.StatusInternalServerError, err
					}
				} else {
					// 所持しているアイテムの場合は所持数更新
					ui.Count += int(item.ItemCount)
					if err := u.userItemRepository.Update(ctx, ui, []string{
						models.UserItemColumns.Count,
						models.UserItemColumns.UpdatedAt,
					}); err != nil {
						return http.StatusInternalServerError, err
					}
				}
			}

			return nil, nil
		}); err != nil {
			return http.StatusInternalServerError, err
		}
	}

	// userの最終ログイン日時更新 (別システムでやる想定でいいかも)
	if u.userRepository.Update(ctx, &models.User{
		ID:          params.UserID,
		LastLoginAt: params.RequestedAt,
	}, []string{
		models.UserColumns.LastLoginAt,
		models.UserColumns.UpdatedAt,
	}); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
