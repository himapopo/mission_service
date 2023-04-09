package repository

import (
	"context"
	"mission_service/models"
)

type UserItemRepository interface {
	Create(context.Context, *models.UserItem) error
	Update(context.Context, *models.UserItem, []string) error
	FetchByItemIDAndUserID(context.Context, int64, int64) (*models.UserItem, error)
	FetchByUserID(context.Context, int64) ([]*models.UserItem, error)
}
