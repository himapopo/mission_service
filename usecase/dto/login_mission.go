package dto

import "time"

type LoginMissionRequest struct {
	UserID      int64     `json:"user_id"`
	RequestedAt time.Time `json:"requested_at"`
}
