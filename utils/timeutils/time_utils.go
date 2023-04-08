package timeutils

import "time"

func DailyMissionResetTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 19, 0, 0, 0, time.UTC)
}
