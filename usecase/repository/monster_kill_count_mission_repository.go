package repository

import (
	"context"
	"mission_service/models"
)

type MonsterKillCountMissionRepository interface {
	FetchNotCompletedByUserIDAndKillCount(context.Context, int64, int64) ([]*models.MonsterKillCountMission, error)
}
