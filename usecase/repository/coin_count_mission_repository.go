package repository

import (
	"context"
	"mission_service/models"
)

type CoinCountMissionRepository interface {
	FetchNotCompletedByUserIDAndCoinCount(context.Context, int64, int64) ([]*models.CoinCountMission, error)
}
