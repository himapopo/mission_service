package mission

import (
	"context"
	"mission_service/models"
	"mission_service/usecase/repository"
	"time"

	"github.com/volatiletech/null/v8"
)

type NormalMissionUsecase interface {
	MonsterKillMission(context.Context, int64, int64, int64, time.Time) error
	MonsterLevelUpMission(context.Context, int64, int64, int, time.Time) error
	MonsterLevelUpCountMission(context.Context, int64, time.Time) error
}

type normalMissionUsecase struct {
	userMissionRepository                repository.UserMissionRepository
	userMissionProgressRepository        repository.UserMissionProgressRepository
	userMonsterRepository                repository.UserMonsterRepository
	monsterKillMissionRepository         repository.MonsterKillMissionRepository
	monsterLevelUpMissionRepository      repository.MonsterLevelUpMissionRepository
	monsterLevelUpCountMissionRepository repository.MonsterLevelUpCountMissionRepository
	missionRewardUsecase                 MissionRewardUsecase
	missionReleaseUsecase                MissionReleaseUsecase
}

func NewNormalMissionUsecase(
	userMissionRepository repository.UserMissionRepository,
	userMissionProgressRepository repository.UserMissionProgressRepository,
	userMonsterRepository repository.UserMonsterRepository,
	monsterKillMissionRepository repository.MonsterKillMissionRepository,
	monsterLevelUpMissionRepository repository.MonsterLevelUpMissionRepository,
	monsterLevelUpCountMissionRepository repository.MonsterLevelUpCountMissionRepository,
	missionRewardUsecase MissionRewardUsecase,
	missionReleaseUsecase MissionReleaseUsecase,
) normalMissionUsecase {
	return normalMissionUsecase{
		userMissionRepository:                userMissionRepository,
		userMissionProgressRepository:        userMissionProgressRepository,
		userMonsterRepository:                userMonsterRepository,
		monsterKillMissionRepository:         monsterKillMissionRepository,
		monsterLevelUpMissionRepository:      monsterLevelUpMissionRepository,
		monsterLevelUpCountMissionRepository: monsterLevelUpCountMissionRepository,
		missionRewardUsecase:                 missionRewardUsecase,
		missionReleaseUsecase:                missionReleaseUsecase,
	}
}

// 一定レベル以上のモンスター獲得ミッション
func (u normalMissionUsecase) MonsterLevelUpCountMission(ctx context.Context, userID int64, requestedAt time.Time) error {
	mlucms, err := u.monsterLevelUpCountMissionRepository.FetchNotCompletedByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if len(mlucms) == 0 {
		return nil
	}

	ms, err := u.userMonsterRepository.FetchByUserID(ctx, userID)
	if err != nil {
		return err
	}

	for _, mlucm := range mlucms {
		targetMonsters := []*models.UserMonster{}
		for _, m := range ms {
			if m.Level >= mlucm.Level {
				targetMonsters = append(targetMonsters, m)
			}
		}

		ump := mlucm.R.Mission.R.UserMissions[0].R.UserMissionProgresses[0]
		ump.ProgressValue = int64(len(targetMonsters))
		ump.LastProgressUpdatedAt = requestedAt
		// ミッションの進捗更新
		if err := u.userMissionProgressRepository.Update(ctx, ump, []string{
			models.UserMissionProgressColumns.ProgressValue,
			models.UserMissionProgressColumns.LastProgressUpdatedAt,
			models.UserMissionProgressColumns.UpdatedAt,
		}); err != nil {
			return err
		}

		// ミッション未達成
		if mlucm.MonsterCount > int64(len(targetMonsters)) {
			continue
		}

		// ミッション達成日時更新
		um := mlucm.R.Mission.R.UserMissions[0]
		um.CompletedAt = null.TimeFrom(requestedAt)
		if err := u.userMissionRepository.Update(ctx, um, []string{
			models.UserMissionColumns.CompletedAt,
			models.UserMissionColumns.UpdatedAt,
		}); err != nil {
			return err
		}

		mission := mlucm.R.Mission

		// ミッション解放
		if len(mission.R.CompleteMissionMissionReleases) != 0 {
			if err := u.missionReleaseUsecase.MissionRelease(ctx, userID, mission.R.CompleteMissionMissionReleases); err != nil {
				return err
			}
		}

		// ミッション報酬獲得
		if err := u.missionRewardUsecase.ObtainRewards(ctx, userID, mission, requestedAt); err != nil {
			return err
		}

	}

	return nil
}

// 特定のモンスターレベルアップミッション
func (u normalMissionUsecase) MonsterLevelUpMission(ctx context.Context, userID, userMonsterID int64, amount int, requestedAt time.Time) error {
	// モンスターのレベルアップ
	m, err := u.userMonsterRepository.Fetch(ctx, userMonsterID)
	if err != nil {
		return err
	}
	m.Level += amount
	if err := u.userMonsterRepository.Update(ctx, m, []string{
		models.UserMonsterColumns.Level,
		models.UserMonsterColumns.UpdatedAt,
	}); err != nil {
		return err
	}

	// ミッション達成チェック
	mlums, err := u.monsterLevelUpMissionRepository.FetchNotCompletedByUserIDAndMonsterID(ctx, userID, m.R.Monster.ID)
	if err != nil {
		return err
	}

	for _, mlum := range mlums {
		ump := mlum.R.Mission.R.UserMissions[0].R.UserMissionProgresses[0]
		ump.ProgressValue = int64(m.Level)
		ump.LastProgressUpdatedAt = requestedAt
		// ミッションの進捗更新
		if err := u.userMissionProgressRepository.Update(ctx, ump, []string{
			models.UserMissionProgressColumns.ProgressValue,
			models.UserMissionProgressColumns.LastProgressUpdatedAt,
			models.UserMissionProgressColumns.UpdatedAt,
		}); err != nil {
			return err
		}

		// ミッション未達成
		if mlum.Level > m.Level {
			continue
		}

		// ミッション達成日時更新
		um := mlum.R.Mission.R.UserMissions[0]
		um.CompletedAt = null.TimeFrom(requestedAt)
		if err := u.userMissionRepository.Update(ctx, um, []string{
			models.UserMissionColumns.CompletedAt,
			models.UserMissionColumns.UpdatedAt,
		}); err != nil {
			return err
		}

		mission := mlum.R.Mission

		// ミッション解放
		if len(mission.R.CompleteMissionMissionReleases) != 0 {
			if err := u.missionReleaseUsecase.MissionRelease(ctx, userID, mission.R.CompleteMissionMissionReleases); err != nil {
				return err
			}
		}

		// ミッション報酬獲得
		if err := u.missionRewardUsecase.ObtainRewards(ctx, userID, mission, requestedAt); err != nil {
			return err
		}

	}

	return nil
}

// 特定のモンスター討伐ミッション
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

		// ミッション未達成
		if mkm.MonsterCount > ump.ProgressValue {
			return nil
		}

		// ミッション達成日時更新
		um := mkm.R.Mission.R.UserMissions[0]
		um.CompletedAt = null.TimeFrom(requestedAt)
		if err := u.userMissionRepository.Update(ctx, um, []string{
			models.UserMissionColumns.CompletedAt,
			models.UserMissionColumns.UpdatedAt,
		}); err != nil {
			return err
		}

		mission := mkm.R.Mission

		// ミッション解放
		if len(mission.R.CompleteMissionMissionReleases) != 0 {
			if err := u.missionReleaseUsecase.MissionRelease(ctx, userID, mission.R.CompleteMissionMissionReleases); err != nil {
				return err
			}
		}

		// ミッション報酬獲得
		if err := u.missionRewardUsecase.ObtainRewards(ctx, userID, mission, requestedAt); err != nil {
			return err
		}

	}

	return nil
}
