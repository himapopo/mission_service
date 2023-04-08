package database

import (
	"context"
	"fmt"
	"mission_service/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type coinCountMissionRepostitory struct {
	dbUtil
}

func NewCoinCountMissionRepostitory(dbUtil dbUtil) coinCountMissionRepostitory {
	return coinCountMissionRepostitory{
		dbUtil: dbUtil,
	}
}

func (r coinCountMissionRepostitory) FetchNotCompletedByUserIDAndCoinCount(ctx context.Context, userID, coinCount int64) ([]*models.CoinCountMission, error) {
	results, err := models.CoinCountMissions(
		qm.InnerJoin(fmt.Sprintf("%s on %s.%s = %s.%s",
			models.TableNames.Missions,
			models.TableNames.Missions,
			models.MissionColumns.ID,
			models.TableNames.CoinCountMissions,
			models.CoinCountMissionColumns.MissionID,
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
		models.CoinCountMissionWhere.CoinCount.LTE(coinCount),
		models.UserMissionWhere.UserID.EQ(userID),
		models.UserMissionWhere.CompletedAt.IsNull(),
		qm.Load(
			qm.Rels(
				models.CoinCountMissionRels.Mission,
				models.MissionRels.UserMissions,
			),
			models.UserMissionWhere.UserID.EQ(userID),
		),
		qm.Load(
			qm.Rels(
				models.CoinCountMissionRels.Mission,
				models.MissionRels.UserMissions,
				models.UserMissionRels.User,
			),
		),
		qm.Load(
			qm.Rels(
				models.CoinCountMissionRels.Mission,
				models.MissionRels.MissionRewardCoins,
			),
		),
		qm.Load(
			qm.Rels(
				models.CoinCountMissionRels.Mission,
				models.MissionRels.MissionRewardItems,
			),
		),
	).All(ctx, r.GetDao(ctx))
	return results, r.Error(err)
}
