package repository

import (
	"context"
	"mission_service/models"
)

type UserMissionProgressRepository interface {
	Update(context.Context, *models.UserMissionProgress, []string) error
	Create(context.Context, *models.UserMissionProgress) error
}
