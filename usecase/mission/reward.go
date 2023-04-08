package mission

import (
	"context"
	"mission_service/models"
	"mission_service/usecase/repository"
)

type MissionRewardUsecase interface {
	ObtainRewards(context.Context, *models.User, *models.Mission) error
}

type missionRewardUsecase struct {
	userRepository     repository.UserRepository
	userItemRepository repository.UserItemRepository
}

func NewMissionRewardUsecase(
	userRepository repository.UserRepository,
	userItemRepository repository.UserItemRepository,
) missionRewardUsecase {
	return missionRewardUsecase{
		userRepository:     userRepository,
		userItemRepository: userItemRepository,
	}
}

func (u missionRewardUsecase) ObtainRewards(ctx context.Context, user *models.User, mission *models.Mission) error {
	rewardCoins := mission.R.MissionRewardCoins
	rewardItems := mission.R.MissionRewardItems

	// コイン報酬加算
	if len(rewardCoins) != 0 {
		user.CoinCount += rewardCoins[0].CoinCount
		if err := u.userRepository.Update(ctx, user, []string{
			models.UserColumns.CoinCount,
			models.UserColumns.UpdatedAt,
		}); err != nil {
			return err
		}
	}

	// アイテム報酬加算
	for _, item := range rewardItems {
		ui, err := u.userItemRepository.FetchByItemIDAndUserID(ctx, user.ID, item.ItemID)
		if err != nil {
			return err
		}
		if ui == nil {
			// 所持していないアイテムの場合はレコード作成
			if err := u.userItemRepository.Create(ctx, &models.UserItem{
				UserID: user.ID,
				ItemID: item.ItemID,
				Count:  int(item.ItemCount),
			}); err != nil {
				return err
			}
		} else {
			// 所持しているアイテムの場合は所持数更新
			ui.Count += int(item.ItemCount)
			if err := u.userItemRepository.Update(ctx, ui, []string{
				models.UserItemColumns.Count,
				models.UserItemColumns.UpdatedAt,
			}); err != nil {
				return err
			}
		}
	}
	return nil
}
