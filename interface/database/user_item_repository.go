package database

import (
	"context"
	"fmt"
	"mission_service/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type userItemRepostitory struct {
	dbUtil
}

func NewUserItemRepostitory(dbUtil dbUtil) userItemRepostitory {
	return userItemRepostitory{
		dbUtil: dbUtil,
	}
}

func (r userItemRepostitory) FetchByItemIDAndUserID(ctx context.Context, userID, itemID int64) (*models.UserItem, error) {
	result, err := models.UserItems(
		models.UserItemWhere.UserID.EQ(userID),
		models.UserItemWhere.ItemID.EQ(itemID),
	).One(ctx, boil.GetContextDB())
	return result, r.Error(err)
}

func (r userItemRepostitory) Create(ctx context.Context, m *models.UserItem) error {
	return r.Error(m.Insert(ctx, boil.GetContextDB(), boil.Infer()))
}

func (r userItemRepostitory) Update(ctx context.Context, m *models.UserItem, updateColumns []string) error {
	cnt, err := m.Update(ctx, boil.GetContextDB(), boil.Whitelist(updateColumns...))
	if cnt == 0 {
		return fmt.Errorf("user item update cnt = %d", cnt)
	}
	return r.Error(err)
}
