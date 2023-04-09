package repository

import (
	"context"
	"mission_service/models"
)

type MonsterLevelUpMissionRepository interface {
	FetchNotCompletedByUserIDAndMonsterID(context.Context, int64, int64) ([]*models.MonsterLevelUpMission, error)
}
