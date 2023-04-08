package repository

import (
	"context"
	"mission_service/models"
)

type LoginMissionRepository interface {
	FetchByUserIDAndLoginCount(context.Context, int64, int64) (*models.LoginMission, error)
}
