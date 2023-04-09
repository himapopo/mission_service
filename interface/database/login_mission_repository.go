package database

import (
	"context"
	"fmt"
	"mission_service/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type loginMissionRepository struct {
	dbUtil
}

func NewLoginMissionRepository(dbUtil dbUtil) loginMissionRepository {
	return loginMissionRepository{
		dbUtil: dbUtil,
	}
}

func (r loginMissionRepository) FetchDailyByUserID(ctx context.Context, userID int64) (*models.LoginMission, error) {
	result, err := models.LoginMissions(
		qm.InnerJoin(fmt.Sprintf("%s on %s.%s = %s.%s",
			models.TableNames.Missions,
			models.TableNames.Missions,
			models.MissionColumns.ID,
			models.TableNames.LoginMissions,
			models.LoginMissionColumns.MissionID,
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
		models.MissionWhere.MissionType.EQ("daily"),
		models.UserMissionWhere.UserID.EQ(userID),
		qm.Load(
			qm.Rels(
				models.LoginMissionRels.Mission,
				models.MissionRels.UserMissions,
			),
			models.UserMissionWhere.UserID.EQ(userID),
		),
		qm.Load(
			qm.Rels(
				models.LoginMissionRels.Mission,
				models.MissionRels.UserMissions,
				models.UserMissionRels.User,
			),
		),
		qm.Load(
			qm.Rels(
				models.LoginMissionRels.Mission,
				models.MissionRels.UserMissions,
				models.UserMissionRels.UserMissionProgresses,
			),
		),
		qm.Load(
			qm.Rels(
				models.LoginMissionRels.Mission,
				models.MissionRels.MissionRewardCoins,
			),
		),
		qm.Load(
			qm.Rels(
				models.LoginMissionRels.Mission,
				models.MissionRels.MissionRewardItems,
			),
		),
		qm.Load(
			qm.Rels(
				models.LoginMissionRels.Mission,
				models.MissionRels.CompleteMissionMissionReleases,
				models.MissionReleaseRels.ReleaseMission,
			),
		),
	).One(ctx, r.GetDao(ctx))
	return result, r.Error(err)
}
