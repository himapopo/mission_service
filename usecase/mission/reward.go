package mission

import (
	"context"
	"mission_service/models"
	"mission_service/usecase/repository"
	"time"

	"github.com/volatiletech/null/v8"
)

type MissionRewardUsecase interface {
	ObtainRewards(context.Context, int64, *models.Mission, time.Time) error
}

type missionRewardUsecase struct {
	userRepository                repository.UserRepository
	userItemRepository            repository.UserItemRepository
	coinCountMissionRepository    repository.CoinCountMissionRepository
	userMissionRepository         repository.UserMissionRepository
	userMissionProgressRepository repository.UserMissionProgressRepository
	getItemMissionRepository      repository.GetItemMissionRepository
	missionReleaseUsecase         MissionReleaseUsecase
}

func NewMissionRewardUsecase(
	userRepository repository.UserRepository,
	userItemRepository repository.UserItemRepository,
	coinCountMissionRepository repository.CoinCountMissionRepository,
	userMissionRepository repository.UserMissionRepository,
	userMissionProgressRepository repository.UserMissionProgressRepository,
	getItemMissionRepository repository.GetItemMissionRepository,
	missionReleaseUsecase MissionReleaseUsecase,
) missionRewardUsecase {
	return missionRewardUsecase{
		userRepository:                userRepository,
		userItemRepository:            userItemRepository,
		coinCountMissionRepository:    coinCountMissionRepository,
		userMissionRepository:         userMissionRepository,
		userMissionProgressRepository: userMissionProgressRepository,
		getItemMissionRepository:      getItemMissionRepository,
		missionReleaseUsecase:         missionReleaseUsecase,
	}

}

func (u missionRewardUsecase) ObtainRewards(ctx context.Context, userID int64, mission *models.Mission, requestedAt time.Time) error {
	rewardCoins := mission.R.MissionRewardCoins
	rewardItems := mission.R.MissionRewardItems
	user, err := u.userRepository.Fetch(ctx, userID)
	if err != nil {
		return err
	}

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
		ui, err := u.userItemRepository.FetchByItemIDAndUserID(ctx, userID, item.ItemID)
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

	// コイン獲得枚数ミッション達成チェック
	if err := u.coinCountMission(ctx, userID, requestedAt); err != nil {
		return err
	}

	// アイテム獲得ミッション達成チェック
	if err := u.getItemMission(ctx, userID, requestedAt); err != nil {
		return err
	}
	return nil
}

// コイン獲得枚数ミッション
func (u missionRewardUsecase) coinCountMission(ctx context.Context, userID int64, requestedAt time.Time) error {
	ccms, err := u.coinCountMissionRepository.FetchNotCompletedByUserID(ctx, userID)
	if err != nil {
		return err
	}
	user, err := u.userRepository.Fetch(ctx, userID)
	if err != nil {
		return err
	}
	for _, ccm := range ccms {
		// ミッションの進捗更新
		ump := ccm.R.Mission.R.UserMissions[0].R.UserMissionProgresses[0]
		ump.ProgressValue = user.CoinCount
		ump.LastProgressUpdatedAt = requestedAt
		if err := u.userMissionProgressRepository.Update(ctx, ump, []string{
			models.UserMissionProgressColumns.ProgressValue,
			models.UserMissionProgressColumns.LastProgressUpdatedAt,
			models.UserMissionProgressColumns.UpdatedAt,
		}); err != nil {
			return err
		}

		// ミッションの未達成
		if ccm.CoinCount > user.CoinCount {
			continue
		}

		// ミッションの達成日時更新
		um := ccm.R.Mission.R.UserMissions[0]
		um.CompletedAt = null.TimeFrom(requestedAt)
		if err := u.userMissionRepository.Update(ctx, um, []string{
			models.UserMissionColumns.CompletedAt,
			models.UserMissionColumns.UpdatedAt,
		}); err != nil {
			return err
		}

		mission := ccm.R.Mission

		// ミッション解放
		if len(mission.R.CompleteMissionMissionReleases) != 0 {
			if err := u.missionReleaseUsecase.MissionRelease(ctx, userID, mission.R.CompleteMissionMissionReleases); err != nil {
				return err
			}
		}

		// ミッション報酬獲得
		if err := u.ObtainRewards(ctx, userID, mission, requestedAt); err != nil {
			return err
		}

	}

	return nil
}

// アイテム獲得ミッション
func (u missionRewardUsecase) getItemMission(ctx context.Context, userID int64, requestedAt time.Time) error {
	gims, err := u.getItemMissionRepository.FetchNotCompletedByUserID(ctx, userID)
	if err != nil {
		return err
	}
	uis, err := u.userItemRepository.FetchByUserID(ctx, userID)
	if err != nil {
		return err
	}
	for _, gim := range gims {
		for _, ui := range uis {
			if gim.ItemID == ui.R.Item.ID {
				// ミッションの進捗更新
				ump := gim.R.Mission.R.UserMissions[0].R.UserMissionProgresses[0]
				ump.ProgressValue = int64(ui.Count)
				ump.LastProgressUpdatedAt = requestedAt
				if err := u.userMissionProgressRepository.Update(ctx, ump, []string{
					models.UserMissionProgressColumns.ProgressValue,
					models.UserMissionProgressColumns.LastProgressUpdatedAt,
					models.UserMissionProgressColumns.UpdatedAt,
				}); err != nil {
					return err
				}
			}

			// ミッション未達成
			if gim.ItemCount > int64(ui.Count) {
				continue
			}

			// ミッションの達成日時更新
			um := gim.R.Mission.R.UserMissions[0]
			um.CompletedAt = null.TimeFrom(requestedAt)
			if err := u.userMissionRepository.Update(ctx, um, []string{
				models.UserMissionColumns.CompletedAt,
				models.UserMissionColumns.UpdatedAt,
			}); err != nil {
				return err
			}

			mission := gim.R.Mission

			// ミッション解放
			if len(mission.R.CompleteMissionMissionReleases) != 0 {
				if err := u.missionReleaseUsecase.MissionRelease(ctx, userID, mission.R.CompleteMissionMissionReleases); err != nil {
					return err
				}
			}

			// ミッション報酬獲得
			if err := u.ObtainRewards(ctx, userID, mission, requestedAt); err != nil {
				return err
			}
		}

	}

	return nil
}
