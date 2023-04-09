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
	CheckCoinCountMission(context.Context, int64, time.Time) error
	MonsterKill(ctx context.Context, params dto.MonsterKillRequest) (int, error)
}

type normalMissionUsecase struct {
	coinCountMissionRepository    repository.CoinCountMissionRepository
	userRepository                repository.UserRepository
	userMissionRepository         repository.UserMissionRepository
	userMissionProgressRepository repository.UserMissionProgressRepository
	monsterKillMissionRepository  repository.MonsterKillMissionRepository
	missionRewardUsecase          MissionRewardUsecase
	weeklyMissionUsecase          WeeklyMissionUsecase
}

func NewNormailMissionUsecase(
	coinCountMissionRepository repository.CoinCountMissionRepository,
	userRepository repository.UserRepository,
	userMissionRepository repository.UserMissionRepository,
	userMissionProgressRepository repository.UserMissionProgressRepository,
	monsterKillMissionRepository repository.MonsterKillMissionRepository,
	missionRewardUsecase MissionRewardUsecase,
	weeklyMissionUsecase WeeklyMissionUsecase,
) normalMissionUsecase {
	return normalMissionUsecase{
		coinCountMissionRepository:    coinCountMissionRepository,
		userRepository:                userRepository,
		userMissionRepository:         userMissionRepository,
		userMissionProgressRepository: userMissionProgressRepository,
		monsterKillMissionRepository:  monsterKillMissionRepository,
		missionRewardUsecase:          missionRewardUsecase,
		weeklyMissionUsecase:          weeklyMissionUsecase,
	}
}

// 特定のモンスター討伐
func (u normalMissionUsecase) MonsterKill(ctx context.Context, params dto.MonsterKillRequest) (int, error) {
	mkm, err := u.monsterKillMissionRepository.FetchNotCompletedByUserIDAndMonsterID(ctx, params.UserID, params.OpponentMonsterID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if err := db.InTx(ctx, func(ctx context.Context) error {
		if mkm != nil {
			ump := mkm.R.Mission.R.UserMissions[0].R.UserMissionProgresses[0]
			ump.ProgressValue += 1
			ump.LastProgressUpdatedAt = params.RequestedAt
			// ミッションの進捗更新
			if err := u.userMissionProgressRepository.Update(ctx, ump, []string{
				models.UserMissionProgressColumns.ProgressValue,
				models.UserMissionProgressColumns.LastProgressUpdatedAt,
				models.UserMissionProgressColumns.UpdatedAt,
			}); err != nil {
				return err
			}

			// ミッション達成時
			if mkm.MonsterCount <= ump.ProgressValue {
				// ミッション達成日時更新
				um := mkm.R.Mission.R.UserMissions[0]
				um.CompletedAt = null.TimeFrom(params.RequestedAt)
				if err := u.userMissionRepository.Update(ctx, um, []string{
					models.UserMissionColumns.CompletedAt,
					models.UserMissionColumns.UpdatedAt,
				}); err != nil {
					return err
				}
				// ミッション報酬獲得
				if err := u.missionRewardUsecase.ObtainRewards(ctx, params.UserID, mkm.R.Mission); err != nil {
					return err
				}
			}
		}

		// 任意のモンスター討伐数ミッション達成チェック
		if err := u.weeklyMissionUsecase.CheckMonsterKillCountMission(ctx, params.UserID, params.RequestedAt); err != nil {
			return err
		}

		// コイン獲得枚数ミッション達成チェック
		if err := u.CheckCoinCountMission(ctx, params.UserID, params.RequestedAt); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// ユーザーの所有コイン数から、達成済みのコイン獲得枚数ミッションがないかチェック
func (u normalMissionUsecase) CheckCoinCountMission(ctx context.Context, userID int64, requestedAt time.Time) error {
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

		// ミッションの達成条件に満たない場合はスキップ
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

		// ミッション報酬獲得
		if err := u.missionRewardUsecase.ObtainRewards(ctx, userID, ccm.R.Mission); err != nil {
			return err
		}

	}
	return nil
}
