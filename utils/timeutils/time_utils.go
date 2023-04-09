package timeutils

import "time"

// 受け取った日の午前4時を返す（日本時間）
func DailyMissionResetTime(t time.Time) time.Time {
	t = t.In(AsiaTokyoLocaion())
	return time.Date(t.Year(), t.Month(), t.Day(), 4, 0, 0, 0, AsiaTokyoLocaion())
}

// 受け取った日の週の月曜日0時を返す（日本時間）
func WeeklyMissionResetTime(t time.Time) time.Time {
	t = t.In(AsiaTokyoLocaion())
	wd := int(t.Weekday())
	daysUntilMonday := 1 - wd
	if daysUntilMonday > 0 {
		daysUntilMonday -= 7
	}
	return time.Date(t.Year(), t.Month(), t.Day()+daysUntilMonday, 0, 0, 0, 0, AsiaTokyoLocaion())
}

func AsiaTokyoLocaion() *time.Location {
	return time.FixedZone("Asia/Tokyo", 9*60*60)
}