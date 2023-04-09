package mission

import (
	"context"
	"mission_service/models"
	"mission_service/usecase/repository"
)

type MissionReleaseUsecase interface {
	MissionRelease(context.Context, int64, []*models.MissionRelease) error
}

type missionReleaseUsecase struct {
	userMissionRepository         repository.UserMissionRepository
	userMissionProgressRepository repository.UserMissionProgressRepository
}

func NewMissionReleaseUsecase(
	userMissionRepository repository.UserMissionRepository,
	userMissionProgressRepository repository.UserMissionProgressRepository,
) missionReleaseUsecase {
	return missionReleaseUsecase{
		userMissionRepository:         userMissionRepository,
		userMissionProgressRepository: userMissionProgressRepository,
	}
}

func (u missionReleaseUsecase) MissionRelease(ctx context.Context, userID int64, releases []*models.MissionRelease) error {
	ums, err := u.userMissionRepository.FetchByUserID(ctx, userID)
	if err != nil {
		return err
	}
	for _, release := range releases {
		exists := false
		for _, um := range ums {
			if release.ReleaseMissionID == um.MissionID {
				exists = true
				break
			}
		}

		if exists {
			continue
		}

		um := &models.UserMission{
			UserID:    userID,
			MissionID: release.ReleaseMissionID,
		}
		// ユーザーミッション作成
		if err := u.userMissionRepository.Create(ctx, um); err != nil {
			return err
		}

		// ミッション進捗管理作成
		if err := u.userMissionProgressRepository.Create(ctx, &models.UserMissionProgress{
			UserMissionID: um.ID,
		}); err != nil {
			return err
		}
	}

	return nil
}
