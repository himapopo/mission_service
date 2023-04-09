package database

import (
	"context"
	"fmt"
	"mission_service/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userMonsterRepository struct {
	dbUtil
}

func NewuUserMonsterRepository(dbUtil dbUtil) userMonsterRepository {
	return userMonsterRepository{
		dbUtil: dbUtil,
	}
}

func (r userMonsterRepository) Update(ctx context.Context, m *models.UserMonster, updateColumns []string) error {
	cnt, err := m.Update(ctx, r.GetDao(ctx), boil.Whitelist(updateColumns...))
	if cnt == 0 {
		return fmt.Errorf("user monster update cnt = %d", cnt)
	}
	return r.Error(err)
}

func (r userMonsterRepository) Fetch(ctx context.Context, id int64) (*models.UserMonster, error) {
	result, err := models.UserMonsters(
		models.UserMonsterWhere.ID.EQ(id),
		qm.Load(models.UserMonsterRels.Monster),
	).One(ctx, r.GetDao(ctx))
	return result, r.Error(err)
}

func (r userMonsterRepository) FetchByUserID(ctx context.Context, userID int64) ([]*models.UserMonster, error) {
	results, err := models.UserMonsters(
		models.UserMonsterWhere.UserID.EQ(userID),
		qm.Load(models.UserMonsterRels.Monster),
	).All(ctx, r.GetDao(ctx))
	return results, r.Error(err)
}
