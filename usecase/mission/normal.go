package mission

import (
	"context"
	"mission_service/models"
	"mission_service/usecase/repository"
	"time"

	"github.com/volatiletech/null/v8"
)

type NormailMissionUsecase interface {
	CheckCoinCountMission(context.Context, *models.User, time.Time) error
}

type normailMissionUsecase struct {
	coinCountMissionRepository repository.CoinCountMissionRepository
	userMissionRepository      repository.UserMissionRepository
	missionRewardUsecase       MissionRewardUsecase
}

func NewNormailMissionUsecase(
	coinCountMissionRepository repository.CoinCountMissionRepository,
	userMissionRepository repository.UserMissionRepository,
	missionRewardUsecase MissionRewardUsecase,
) normailMissionUsecase {
	return normailMissionUsecase{
		coinCountMissionRepository: coinCountMissionRepository,
		userMissionRepository:      userMissionRepository,
		missionRewardUsecase:       missionRewardUsecase,
	}
}

// ユーザーの所有コイン数から、達成済みのコイン獲得枚数ミッションがないかチェック
func (u normailMissionUsecase) CheckCoinCountMission(ctx context.Context, user *models.User, completedAt time.Time) error {
	ccms, err := u.coinCountMissionRepository.FetchNotCompletedByUserIDAndCoinCount(ctx, user.ID, user.CoinCount)
	if err != nil {
		return err
	}
	for _, m := range ccms {
		// ミッションの達成日時更新
		um := m.R.Mission.R.UserMissions[0]
		um.CompletedAt = null.TimeFrom(completedAt)
		if err := u.userMissionRepository.Update(ctx, um, []string{
			models.UserMissionColumns.CompletedAt,
			models.UserMissionColumns.UpdatedAt,
		}); err != nil {
			return err
		}

		// ミッション報酬加算
		if err := u.missionRewardUsecase.ObtainRewards(ctx, um.R.User, m.R.Mission); err != nil {
			return err
		}

	}
	return nil
}
