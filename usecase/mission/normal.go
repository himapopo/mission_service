package mission

import (
	"context"
	"mission_service/infrastructure/db"
	"mission_service/models"
	"mission_service/usecase/dto"
	"mission_service/usecase/repository"
	"net/http"
	"time"

	"github.com/volatiletech/null/v8"
)

type NormalMissionUsecase interface {
	CheckCoinCountMission(context.Context, *models.User, time.Time) error
	MonsterKill(ctx context.Context, params dto.MonsterKillMissionRequest) (int, error)
}

type normalMissionUsecase struct {
	coinCountMissionRepository   repository.CoinCountMissionRepository
	userMissionRepository        repository.UserMissionRepository
	monsterKillMissionRepository repository.MonsterKillMissionRepository
	missionRewardUsecase         MissionRewardUsecase
	weeklyMissionUsecase         WeeklyMissionUsecase
}

func NewNormailMissionUsecase(
	coinCountMissionRepository repository.CoinCountMissionRepository,
	userMissionRepository repository.UserMissionRepository,
	monsterKillMissionRepository repository.MonsterKillMissionRepository,
	missionRewardUsecase MissionRewardUsecase,
	weeklyMissionUsecase WeeklyMissionUsecase,
) normalMissionUsecase {
	return normalMissionUsecase{
		coinCountMissionRepository:   coinCountMissionRepository,
		userMissionRepository:        userMissionRepository,
		monsterKillMissionRepository: monsterKillMissionRepository,
		missionRewardUsecase:         missionRewardUsecase,
		weeklyMissionUsecase:         weeklyMissionUsecase,
	}
}

func (u normalMissionUsecase) MonsterKill(ctx context.Context, params dto.MonsterKillMissionRequest) (int, error) {

	// 前回ミッション達成日時が今日の04:00以前の場合はミッション報酬獲得
	if err := db.InTx(ctx, func(ctx context.Context) error {

		// ログインミッション報酬獲得
		// if err := u.missionRewardUsecase.ObtainRewards(ctx, user, lm.R.Mission); err != nil {
		// 	return err
		// }

		// // コイン獲得枚数ミッション達成チェック
		// if err := u.CheckCoinCountMission(ctx, user, params.RequestedAt); err != nil {
		// 	return err
		// }

		return nil
	}); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// ユーザーの所有コイン数から、達成済みのコイン獲得枚数ミッションがないかチェック
func (u normalMissionUsecase) CheckCoinCountMission(ctx context.Context, user *models.User, completedAt time.Time) error {
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
