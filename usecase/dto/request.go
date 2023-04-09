package dto

import "time"

type LoginRequest struct {
	UserID      int64     `json:"user_id"`
	RequestedAt time.Time `json:"requested_at"`
}

type MonsterKillRequest struct {
	UserID            int64     `json:"user_id"`
	UserMonsterID     int64     `json:"user_monster_id"`
	OpponentMonsterID int64     `json:"opponent_monster_id"`
	RequestedAt       time.Time `json:"requested_at"`
}

type MonsterLevelUpRequest struct {
	UserID        int64     `json:"user_id"`
	UserMonsterID int64     `json:"user_monster_id"`
	Amount        int       `json:"amount"`
	RequestedAt   time.Time `json:"requested_at"`
}
