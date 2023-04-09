package repository

import (
	"context"
	"mission_service/models"
)

type UserMonsterRepository interface {
	Update(context.Context, *models.UserMonster, []string) error
	Fetch(context.Context, int64) (*models.UserMonster, error)
}
