package repository

import (
	"context"
	"mission_service/models"
)

type MonsterKillCountMissionRepository interface {
	FetchWeeklyByUserID(context.Context, int64) ([]*models.MonsterKillCountMission, error)
}
