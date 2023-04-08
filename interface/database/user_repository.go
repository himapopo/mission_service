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
	cnt, err := m.Update(ctx, boil.GetContextDB(), boil.Whitelist(updateColumns...))
	if cnt == 0 {
		return fmt.Errorf("user update cnt = %d", cnt)
	}
	return r.Error(err)
}
