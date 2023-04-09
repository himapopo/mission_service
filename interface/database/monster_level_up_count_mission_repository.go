package database

import (
	"context"
	"fmt"
	"mission_service/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type monsterLevelUpCountMissionRepository struct {
	dbUtil
}

func NewMonsterLevelUpCountMissionRepository(dbUtil dbUtil) monsterLevelUpCountMissionRepository {
	return monsterLevelUpCountMissionRepository{
		dbUtil: dbUtil,
	}
}

func (r monsterLevelUpCountMissionRepository) FetchNotCompletedByUserID(ctx context.Context, userID int64) ([]*models.MonsterLevelUpCountMission, error) {
	results, err := models.MonsterLevelUpCountMissions(
		qm.InnerJoin(fmt.Sprintf("%s on %s.%s = %s.%s",
			models.TableNames.Missions,
			models.TableNames.Missions,
			models.MissionColumns.ID,
			models.TableNames.MonsterLevelUpCountMissions,
			models.MonsterLevelUpCountMissionColumns.MissionID,
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
		models.UserMissionWhere.UserID.EQ(userID),
		models.UserMissionWhere.CompletedAt.IsNull(),
		qm.Load(
			qm.Rels(
				models.MonsterLevelUpCountMissionRels.Mission,
				models.MissionRels.UserMissions,
			),
			models.UserMissionWhere.UserID.EQ(userID),
		),
		qm.Load(
			qm.Rels(
				models.MonsterLevelUpCountMissionRels.Mission,
				models.MissionRels.UserMissions,
				models.UserMissionRels.User,
			),
		),
		qm.Load(
			qm.Rels(
				models.MonsterLevelUpCountMissionRels.Mission,
				models.MissionRels.UserMissions,
				models.UserMissionRels.UserMissionProgresses,
			),
		),
		qm.Load(
			qm.Rels(
				models.MonsterLevelUpCountMissionRels.Mission,
				models.MissionRels.MissionRewardCoins,
			),
		),
		qm.Load(
			qm.Rels(
				models.MonsterLevelUpCountMissionRels.Mission,
				models.MissionRels.MissionRewardItems,
			),
		),
		qm.Load(
			qm.Rels(
				models.MonsterLevelUpCountMissionRels.Mission,
				models.MissionRels.CompleteMissionMissionReleases,
				models.MissionReleaseRels.ReleaseMission,
			),
		),
	).All(ctx, r.GetDao(ctx))
	return results, r.Error(err)
}
