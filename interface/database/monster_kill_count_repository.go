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

func (r monsterKillCountMissionRepostitory) FetchWeeklyByUserID(ctx context.Context, userID int64) ([]*models.MonsterKillCountMission, error) {
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
		models.MissionWhere.MissionType.EQ("weekly"),
		models.UserMissionWhere.UserID.EQ(userID),
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
				models.MissionRels.UserMissions,
				models.UserMissionRels.UserMissionProgresses,
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
