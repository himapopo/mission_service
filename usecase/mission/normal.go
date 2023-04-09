package mission

import (
	"context"
	"mission_service/models"
	"mission_service/usecase/repository"
	"time"

	"github.com/volatiletech/null/v8"
)

type NormalMissionUsecase interface {
	CoinCountMission(context.Context, int64, time.Time) error
	MonsterKillMission(context.Context, int64, int64, int64, time.Time) error
	MonsterLevelUpMission(context.Context, int64, int64, int64, time.Time) error
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

// 特定のモンスターレベルアップ
func (u normalMissionUsecase) MonsterLevelUpMission(ctx context.Context, userID, myMonsterID, amount int64, requestedAt time.Time) error {

	// 	// 一定レベル以上のモンスター獲得ミッション達成チェック

	// 	// コイン獲得枚数ミッション達成チェック
	// 	if err := u.CheckCoinCountMission(ctx, userID, requestedAt); err != nil {
	// 		return err
	// 	}

	return nil
}

// 特定のモンスター討伐
func (u normalMissionUsecase) MonsterKillMission(ctx context.Context, userID, myMonsterID, opponentMonsterID int64, requestedAt time.Time) error {
	mkm, err := u.monsterKillMissionRepository.FetchNotCompletedByUserIDAndMonsterID(ctx, userID, opponentMonsterID)
	if err != nil {
		return err
	}
	if mkm != nil {
		ump := mkm.R.Mission.R.UserMissions[0].R.UserMissionProgresses[0]
		ump.ProgressValue += 1
		ump.LastProgressUpdatedAt = requestedAt
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
			um.CompletedAt = null.TimeFrom(requestedAt)
			if err := u.userMissionRepository.Update(ctx, um, []string{
				models.UserMissionColumns.CompletedAt,
				models.UserMissionColumns.UpdatedAt,
			}); err != nil {
				return err
			}
			// ミッション報酬獲得
			if err := u.missionRewardUsecase.ObtainRewards(ctx, userID, mkm.R.Mission); err != nil {
				return err
			}
		}
	}

	return nil
}

// コイン獲得枚数ミッション達成チェック
func (u normalMissionUsecase) CoinCountMission(ctx context.Context, userID int64, requestedAt time.Time) error {
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
