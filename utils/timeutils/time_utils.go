package timeutils

import "time"

func DailyMissionResetTime(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 19, 0, 0, 0, time.UTC)
}

// TODO: 実装
func WeeklyMissionResetTime(now time.Time) time.Time {
	// wd := now.Weekday()
	return time.Date(now.Year(), now.Month(), now.Day(), 19, 0, 0, 0, time.UTC)
}
