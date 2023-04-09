package database

import (
	"context"
	"fmt"
	"mission_service/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userItemRepository struct {
	dbUtil
}

func NewUserItemRepository(dbUtil dbUtil) userItemRepository {
	return userItemRepository{
		dbUtil: dbUtil,
	}
}

func (r userItemRepository) FetchByItemIDAndUserID(ctx context.Context, userID, itemID int64) (*models.UserItem, error) {
	result, err := models.UserItems(
		models.UserItemWhere.UserID.EQ(userID),
		models.UserItemWhere.ItemID.EQ(itemID),
	).One(ctx, r.GetDao(ctx))
	return result, r.Error(err)
}

func (r userItemRepository) FetchByUserID(ctx context.Context, userID int64) ([]*models.UserItem, error) {
	results, err := models.UserItems(
		models.UserItemWhere.UserID.EQ(userID),
		qm.Load(models.UserItemRels.Item),
	).All(ctx, r.GetDao(ctx))
	return results, r.Error(err)
}

func (r userItemRepository) Create(ctx context.Context, m *models.UserItem) error {
	return r.Error(m.Insert(ctx, r.GetDao(ctx), boil.Infer()))
}

func (r userItemRepository) Update(ctx context.Context, m *models.UserItem, updateColumns []string) error {
	cnt, err := m.Update(ctx, r.GetDao(ctx), boil.Whitelist(updateColumns...))
	if cnt == 0 {
		return fmt.Errorf("user item update cnt = %d", cnt)
	}
	return r.Error(err)
}
