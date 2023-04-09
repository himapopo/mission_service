package timeutils

import "time"

// 引数の日のデイリーミッションリセット日時を返す
func DailyMissionResetTime(t time.Time) time.Time {
	t = t.In(AsiaTokyoLocaion())
	resetTime := time.Date(t.Year(), t.Month(), t.Day(), 4, 0, 0, 0, AsiaTokyoLocaion())
	if t.Before(resetTime) {
		resetTime = resetTime.AddDate(0, 0, -1)
	}
	return resetTime
}

// 引数の日の週のウィークリーミッションリセット日時を返す
func WeeklyMissionResetTime(t time.Time) time.Time {
	t = t.In(AsiaTokyoLocaion())
	wd := int(t.Weekday())
	daysUntilMonday := 1 - wd
	if daysUntilMonday > 0 {
		daysUntilMonday -= 7
	}
	return time.Date(t.Year(), t.Month(), t.Day()+daysUntilMonday, 0, 0, 0, 0, AsiaTokyoLocaion())
}

// JST
func AsiaTokyoLocaion() *time.Location {
	return time.FixedZone("Asia/Tokyo", 9*60*60)
}
