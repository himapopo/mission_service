package database

import (
	"context"
	"fmt"
	"mission_service/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type getItemMissionRepository struct {
	dbUtil
}

func NewGetItemMissionRepository(dbUtil dbUtil) getItemMissionRepository {
	return getItemMissionRepository{
		dbUtil: dbUtil,
	}
}

func (r getItemMissionRepository) FetchNotCompletedByUserID(ctx context.Context, userID int64) ([]*models.GetItemMission, error) {
	results, err := models.GetItemMissions(
		qm.InnerJoin(fmt.Sprintf("%s on %s.%s = %s.%s",
			models.TableNames.Missions,
			models.TableNames.Missions,
			models.MissionColumns.ID,
			models.TableNames.GetItemMissions,
			models.GetItemMissionColumns.MissionID,
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
				models.GetItemMissionRels.Mission,
				models.MissionRels.UserMissions,
			),
			models.UserMissionWhere.UserID.EQ(userID),
		),
		qm.Load(
			qm.Rels(
				models.GetItemMissionRels.Mission,
				models.MissionRels.UserMissions,
				models.UserMissionRels.User,
			),
		),
		qm.Load(
			qm.Rels(
				models.GetItemMissionRels.Mission,
				models.MissionRels.UserMissions,
				models.UserMissionRels.UserMissionProgresses,
			),
		),
		qm.Load(
			qm.Rels(
				models.GetItemMissionRels.Mission,
				models.MissionRels.MissionRewardCoins,
			),
		),
		qm.Load(
			qm.Rels(
				models.GetItemMissionRels.Mission,
				models.MissionRels.MissionRewardItems,
			),
		),
		qm.Load(
			qm.Rels(
				models.GetItemMissionRels.Mission,
				models.MissionRels.CompleteMissionMissionReleases,
				models.MissionReleaseRels.ReleaseMission,
			),
		),
	).All(ctx, r.GetDao(ctx))
	return results, r.Error(err)
}
