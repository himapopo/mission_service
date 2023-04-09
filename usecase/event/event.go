package event

import (
	"context"
	"mission_service/infrastructure/db"
	"mission_service/usecase/dto"
	"mission_service/usecase/mission"
	"net/http"
)

type EventUsecase interface {
	Login(ctx context.Context, params dto.LoginRequest) (int, error)
	MonsterKill(ctx context.Context, params dto.MonsterKillRequest) (int, error)
	MonsterLevelUp(ctx context.Context, params dto.MonsterLevelUpRequest) (int, error)
}

type eventUsecase struct {
	dailyMissionUsecase  mission.DailyMissionUsecase
	weeklyMissionUsecase mission.WeeklyMissionUsecase
	normalMissionUsecase mission.NormalMissionUsecase
}

func NewEventMissionUsecase(
	dailyMissionUsecase mission.DailyMissionUsecase,
	weeklyMissionUsecase mission.WeeklyMissionUsecase,
	normalMissionUsecase mission.NormalMissionUsecase,
) eventUsecase {
	return eventUsecase{
		dailyMissionUsecase:  dailyMissionUsecase,
		weeklyMissionUsecase: weeklyMissionUsecase,
		normalMissionUsecase: normalMissionUsecase,
	}
}

func (u eventUsecase) Login(ctx context.Context, params dto.LoginRequest) (int, error) {
	if err := db.InTx(ctx, func(ctx context.Context) error {

		// ログインミッション達成チェック
		if err := u.dailyMissionUsecase.LoginMission(ctx, params.UserID, params.RequestedAt); err != nil {
			return err
		}

		// コイン獲得枚数ミッション達成チェック
		if err := u.normalMissionUsecase.CoinCountMission(ctx, params.UserID, params.RequestedAt); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (u eventUsecase) MonsterKill(ctx context.Context, params dto.MonsterKillRequest) (int, error) {
	if err := db.InTx(ctx, func(ctx context.Context) error {

		// 特定のモンスター討伐ミッション達成チェック
		if err := u.normalMissionUsecase.MonsterKillMission(ctx, params.UserID, params.MyMonsterID, params.OpponentMonsterID, params.RequestedAt); err != nil {
			return err
		}

		// 任意のモンスター討伐数ミッション達成チェック
		if err := u.weeklyMissionUsecase.MonsterKillCountMission(ctx, params.UserID, params.RequestedAt); err != nil {
			return err
		}

		// コイン獲得枚数ミッション達成チェック
		if err := u.normalMissionUsecase.CoinCountMission(ctx, params.UserID, params.RequestedAt); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (u eventUsecase) MonsterLevelUp(ctx context.Context, params dto.MonsterLevelUpRequest) (int, error) {
	if err := db.InTx(ctx, func(ctx context.Context) error {

		// 特定のモンスターレベルアップミッション達成チェック
		if err := u.normalMissionUsecase.MonsterLevelUpMission(ctx, params.UserID, params.MyMonsterID, params.Amount, params.RequestedAt); err != nil {
			return err
		}

		// 一定レベル以上のモンスター獲得ミッション達成チェック
		if err := u.normalMissionUsecase.MonsterLevelUpCountMission(ctx, params.UserID, params.RequestedAt); err != nil {
			return err
		}

		// コイン獲得枚数ミッション達成チェック
		if err := u.normalMissionUsecase.CoinCountMission(ctx, params.UserID, params.RequestedAt); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
