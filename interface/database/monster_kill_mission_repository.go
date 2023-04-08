package database

import (
	"context"
	"fmt"
	"mission_service/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type monsterKillMissionRepostitory struct {
	dbUtil
}

func NewMonsterKillMissionRepostitory(dbUtil dbUtil) monsterKillMissionRepostitory {
	return monsterKillMissionRepostitory{
		dbUtil: dbUtil,
	}
}

func (r monsterKillMissionRepostitory) FetchNotCompletedByUserIDAndMonsterID(ctx context.Context, userID, monsterID int64) (*models.MonsterKillMission, error) {
	result, err := models.MonsterKillMissions(
		qm.InnerJoin(fmt.Sprintf("%s on %s.%s = %s.%s",
			models.TableNames.Missions,
			models.TableNames.Missions,
			models.MissionColumns.ID,
			models.TableNames.MonsterKillMissions,
			models.MonsterKillMissionColumns.MissionID,
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
		models.MonsterKillMissionWhere.MonsterID.EQ(monsterID),
		models.UserMissionWhere.UserID.EQ(userID),
		models.UserMissionWhere.CompletedAt.IsNull(),
		qm.Load(
			qm.Rels(
				models.MonsterKillMissionRels.Mission,
				models.MissionRels.UserMissions,
			),
			models.UserMissionWhere.UserID.EQ(userID),
		),
		qm.Load(
			qm.Rels(
				models.MonsterKillMissionRels.Mission,
				models.MissionRels.UserMissions,
				models.UserMissionRels.User,
			),
		),
		qm.Load(
			qm.Rels(
				models.MonsterKillMissionRels.Mission,
				models.MissionRels.UserMissions,
				models.UserMissionRels.UserMissionProgresses,
			),
		),
		qm.Load(
			qm.Rels(
				models.MonsterKillMissionRels.Mission,
				models.MissionRels.MissionRewardCoins,
			),
		),
		qm.Load(
			qm.Rels(
				models.MonsterKillMissionRels.Mission,
				models.MissionRels.MissionRewardItems,
			),
		),
	).One(ctx, r.GetDao(ctx))
	return result, r.Error(err)
}
