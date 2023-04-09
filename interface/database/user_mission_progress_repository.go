package database

import (
	"context"
	"fmt"
	"mission_service/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type userMissionProgressRepository struct {
	dbUtil
}

func NewUserMissionProgressRepository(dbUtil dbUtil) userMissionProgressRepository {
	return userMissionProgressRepository{
		dbUtil: dbUtil,
	}
}

func (r userMissionProgressRepository) Update(ctx context.Context, m *models.UserMissionProgress, updateColumns []string) error {
	cnt, err := m.Update(ctx, r.GetDao(ctx), boil.Whitelist(updateColumns...))
	if cnt == 0 {
		return fmt.Errorf("user mission progress update cnt = %d", cnt)
	}
	return r.Error(err)
}
