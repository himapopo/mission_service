package repository

import (
	"context"
	"mission_service/models"
)

type MonsterLevelUpCountMissionRepository interface {
	FetchNotCompletedByUserID(context.Context, int64) ([]*models.MonsterLevelUpCountMission, error)
}
