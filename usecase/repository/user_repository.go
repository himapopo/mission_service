package repository

import (
	"context"
	"mission_service/models"
)

type UserRepository interface {
	Update(context.Context, *models.User, []string) error
	Fetch(context.Context, int64) (*models.User, error)
}
