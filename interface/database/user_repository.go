package database

import (
	"context"
	"fmt"
	"mission_service/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type userRepostitory struct {
	dbUtil
}

func NewuUserRepostitory(dbUtil dbUtil) userRepostitory {
	return userRepostitory{
		dbUtil: dbUtil,
	}
}

func (r userRepostitory) Update(ctx context.Context, m *models.User, updateColumns []string) error {
	cnt, err := m.Update(ctx, r.GetDao(ctx), boil.Whitelist(updateColumns...))
	if cnt == 0 {
		return fmt.Errorf("user update cnt = %d", cnt)
	}
	return r.Error(err)
}

func (r userRepostitory) Fetch(ctx context.Context, id int64) (*models.User, error) {
	result, err := models.Users(
		models.UserWhere.ID.EQ(id),
	).One(ctx, r.GetDao(ctx))
	return result, r.Error(err)
}
