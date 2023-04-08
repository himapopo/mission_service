package database

import (
	"context"
	"fmt"
	"mission_service/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type userMissionRepostitory struct {
	dbUtil
}

func NewUserMissionRepostitory(dbUtil dbUtil) userMissionRepostitory {
	return userMissionRepostitory{
		dbUtil: dbUtil,
	}
}

func (r userMissionRepostitory) Update(ctx context.Context, m *models.UserMission, updateColumns []string) error {
	cnt, err := m.Update(ctx, r.GetDao(ctx), boil.Whitelist(updateColumns...))
	if cnt == 0 {
		return fmt.Errorf("user mission update cnt = %d", cnt)
	}
	return r.Error(err)
}
