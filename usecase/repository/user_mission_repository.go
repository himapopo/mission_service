package repository

import (
	"context"
	"mission_service/models"
)

type UserMissionRepository interface {
	Update(context.Context, *models.UserMission, []string) error
}
