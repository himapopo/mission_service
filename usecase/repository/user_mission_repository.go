package repository

import (
	"context"
	"mission_service/models"
)

type UserMissionRepository interface {
	Update(context.Context, *models.UserMission, []string) error
	FetchByUserID(context.Context, int64) ([]*models.UserMission, error)
	Create(context.Context, *models.UserMission) error
}
