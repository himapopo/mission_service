package repository

import (
	"context"
	"mission_service/models"
)

type MonsterKillMissionRepository interface {
	FetchNotCompletedByUserIDAndMonsterID(context.Context, int64, int64) (*models.MonsterKillMission, error)
}
