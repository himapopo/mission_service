package timeutils

import (
	"testing"
	"time"
)

func TestDailyMissionResetTime(t *testing.T) {
	tests := []struct {
		name string
		arg  time.Time
		want time.Time
	}{
		{
			name: "4月3日4時を返す",
			arg:  time.Date(2023, 4, 3, 14, 0, 0, 0, time.UTC),
			want: time.Date(2023, 4, 3, 4, 0, 0, 0, AsiaTokyoLocaion()),
		},
		{
			name: "4月5日4時を返す",
			arg:  time.Date(2023, 4, 5, 15, 0, 0, 0, time.UTC),
			want: time.Date(2023, 4, 5, 4, 0, 0, 0, AsiaTokyoLocaion()),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mondy := DailyMissionResetTime(test.arg)
			if !mondy.Equal(test.want) {
				t.Fatalf("result = %v, want = %v", mondy, test.want)
			}
		})
	}
}

func TestWeeklyMissionResetTime(t *testing.T) {
	tests := []struct {
		name string
		arg  time.Time
		want time.Time
	}{
		{
			name: "4月3日0時を返す",
			arg:  time.Date(2023, 4, 5, 20, 0, 0, 0, time.UTC),
			want: time.Date(2023, 4, 3, 0, 0, 0, 0, AsiaTokyoLocaion()),
		},
		{
			name: "3月27日0時を返す",
			arg:  time.Date(2023, 4, 2, 14, 59, 59, 0, time.UTC),
			want: time.Date(2023, 3, 27, 0, 0, 0, 0, AsiaTokyoLocaion()),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mondy := WeeklyMissionResetTime(test.arg)
			if !mondy.Equal(test.want) {
				t.Fatalf("result = %v, want = %v", mondy, test.want)
			}
		})
	}
}
