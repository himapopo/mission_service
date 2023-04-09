package repository

import (
	"context"
	"mission_service/models"
)

type GetItemMissionRepository interface {
	FetchNotCompletedByUserID(context.Context, int64) ([]*models.GetItemMission, error)
}
