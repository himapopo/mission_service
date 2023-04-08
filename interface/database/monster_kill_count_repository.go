package database

import (
	"context"
	"fmt"
	"mission_service/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type monsterKillCountMissionRepostitory struct {
	dbUtil
}

func NewMonsterKillCountMissionRepostitory(dbUtil dbUtil) monsterKillCountMissionRepostitory {
	return monsterKillCountMissionRepostitory{
		dbUtil: dbUtil,
	}
}

func (r monsterKillCountMissionRepostitory) FetchNotCompletedByUserIDAndKillCount(ctx context.Context, userID, killCount int64) ([]*models.MonsterKillCountMission, error) {
	results, err := models.MonsterKillCountMissions(
		qm.InnerJoin(fmt.Sprintf("%s on %s.%s = %s.%s",
			models.TableNames.Missions,
			models.TableNames.Missions,
			models.MissionColumns.ID,
			models.TableNames.MonsterKillCountMissions,
			models.MonsterKillCountMissionColumns.MissionID,
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
		models.MonsterKillCountMissionWhere.KillCount.LTE(killCount),
		models.UserMissionWhere.UserID.EQ(userID),
		models.UserMissionWhere.CompletedAt.IsNull(),
		qm.Load(
			qm.Rels(
				models.MonsterKillCountMissionRels.Mission,
				models.MissionRels.UserMissions,
			),
			models.UserMissionWhere.UserID.EQ(userID),
		),
		qm.Load(
			qm.Rels(
				models.MonsterKillCountMissionRels.Mission,
				models.MissionRels.UserMissions,
				models.UserMissionRels.User,
			),
		),
		qm.Load(
			qm.Rels(
				models.MonsterKillCountMissionRels.Mission,
				models.MissionRels.MissionRewardCoins,
			),
		),
		qm.Load(
			qm.Rels(
				models.MonsterKillCountMissionRels.Mission,
				models.MissionRels.MissionRewardItems,
			),
		),
	).All(ctx, r.GetDao(ctx))
	return results, r.Error(err)
}
