package dto

import "time"

type LoginMissionRequest struct {
	UserID      int64     `json:"user_id"`
	RequestedAt time.Time `json:"requested_at"`
}

type MonsterKillMissionRequest struct {
	UserID      int64     `json:"user_id"`
	MyMonsterID      int64     `json:"my_monster_id"`
	OpponentMonsterID      int64     `json:"opponent_monster_id"`
	RequestedAt time.Time `json:"requested_at"`
}