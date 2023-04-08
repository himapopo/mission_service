package database

import (
	"context"
	"fmt"
	"mission_service/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type loginMissionRepostitory struct {
	dbUtil
}

func NewLoginMissionRepostitory(dbUtil dbUtil) loginMissionRepostitory {
	return loginMissionRepostitory{
		dbUtil: dbUtil,
	}
}

func (r loginMissionRepostitory) FetchByUserIDAndLoginCount(ctx context.Context, userID, loginCount int64) (*models.LoginMission, error) {
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
		models.LoginMissionWhere.LoginCount.EQ(loginCount),
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
				models.MissionRels.MissionRewardCoins,
			),
		),
		qm.Load(
			qm.Rels(
				models.LoginMissionRels.Mission,
				models.MissionRels.MissionRewardItems,
			),
		),
	).One(ctx, boil.GetContextDB())
	return result, r.Error(err)
}
