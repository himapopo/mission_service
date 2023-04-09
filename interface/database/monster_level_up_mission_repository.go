package database

import (
	"context"
	"fmt"
	"mission_service/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type monsterLevelUpMissionRepository struct {
	dbUtil
}

func NewMonsterLevelUpMissionRepository(dbUtil dbUtil) monsterLevelUpMissionRepository {
	return monsterLevelUpMissionRepository{
		dbUtil: dbUtil,
	}
}

func (r monsterLevelUpMissionRepository) FetchNotCompletedByUserIDAndMonsterID(ctx context.Context, userID, monsterID int64) ([]*models.MonsterLevelUpMission, error) {
	results, err := models.MonsterLevelUpMissions(
		qm.InnerJoin(fmt.Sprintf("%s on %s.%s = %s.%s",
			models.TableNames.Missions,
			models.TableNames.Missions,
			models.MissionColumns.ID,
			models.TableNames.MonsterLevelUpMissions,
			models.MonsterLevelUpMissionColumns.MissionID,
		),
		),
		qm.InnerJoin(fmt.Sprintf("%s on %s.%s = %s.%s",
			models.TableNames.UserMissions,
			models.TableNames.UserMissions,
			models.UserMissionColumns.MissionID,
			models.TableNames.Missions,
			models.MissionColumns.ID,
		),
		),
		models.MissionWhere.IsDeleted.EQ(false),
		models.MonsterLevelUpMissionWhere.MonsterID.EQ(monsterID),
		models.UserMissionWhere.UserID.EQ(userID),
		models.UserMissionWhere.CompletedAt.IsNull(),
		qm.Load(
			qm.Rels(
				models.MonsterLevelUpMissionRels.Mission,
				models.MissionRels.UserMissions,
			),
			models.UserMissionWhere.UserID.EQ(userID),
		),
		qm.Load(
			qm.Rels(
				models.MonsterLevelUpMissionRels.Mission,
				models.MissionRels.UserMissions,
				models.UserMissionRels.User,
			),
		),
		qm.Load(
			qm.Rels(
				models.MonsterLevelUpMissionRels.Mission,
				models.MissionRels.UserMissions,
				models.UserMissionRels.UserMissionProgresses,
			),
		),
		qm.Load(
			qm.Rels(
				models.MonsterLevelUpMissionRels.Mission,
				models.MissionRels.MissionRewardCoins,
			),
		),
		qm.Load(
			qm.Rels(
				models.MonsterLevelUpMissionRels.Mission,
				models.MissionRels.MissionRewardItems,
			),
		),
	).All(ctx, r.GetDao(ctx))
	return results, r.Error(err)
}
