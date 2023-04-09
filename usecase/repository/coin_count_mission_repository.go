package repository

import (
	"context"
	"mission_service/models"
)

type CoinCountMissionRepository interface {
	FetchNotCompletedByUserID(context.Context, int64) ([]*models.CoinCountMission, error)
}
