package utils

import (
	"strings"
	"time"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

func GetZero() time.Time {
	now := time.Now()
	dd, _ := time.ParseDuration("24h")
	tomorrow := now.Add(dd)
	year, month, day := tomorrow.Date()
	zero := time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	return zero
}

//StrToTime 字符串转换时间
func StrToTime(str string) (time.Time, error) {
	t, err := time.ParseInLocation(timeLayout, str, time.Local)
	return t, err
}

//TimeToStr 时间转换字符串
func TimeToStr(t time.Time) string {
	format := t.Format(timeLayout)
	format = strings.Replace(format, "T", " ", -1)
	format = strings.Replace(format, "+", " ", -1)
	return format
}

//UnixToStr 时间戳转换字符串
func UnixToStr(u int64) string {
	dataTimeStr := time.Unix(u, 0).Format(timeLayout)
	return dataTimeStr
}

//UnixToTime 时间戳转换字符串
func UnixToTime(u int64) time.Time {
	dataTime := time.Unix(u, 0)
	return dataTime
}

//StrSubDays 时间相差天数 2007-01-03 00:00:00
func StrSubDays(s1 string, s2 string) (int, error) {
	t1, err := StrToTime(s1)
	if err != nil {
		return 0, err
	}
	t2, err := StrToTime(s2)
	if err != nil {
		return 0, err
	}
	t1 = t1.UTC().Truncate(24 * time.Hour)
	t2 = t2.UTC().Truncate(24 * time.Hour)
	newTime := t1.Sub(t2)
	return int(newTime.Hours() / 24), nil
}

//DiffTime 对比时间
func DiffTime(t1, t2 time.Time) time.Duration {
	t1 = t1.UTC()
	t2 = t2.UTC()
	return t1.Sub(t2)
}

//DiffTime 对比时间
func DiffUnix(u1, u2 int64) time.Duration {
	var t1 time.Time = time.Unix(u1, 0)
	var t2 time.Time = time.Unix(u2, 0)
	t1 = t1.UTC()
	t2 = t2.UTC()
	return t1.Sub(t2)
}
