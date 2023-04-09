package repository

import (
	"context"
	"mission_service/models"
)

type LoginMissionRepository interface {
	FetchDailyByUserID(context.Context, int64) (*models.LoginMission, error)
}
