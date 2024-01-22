package utils

import "time"

var defaultTimeZone, _ = time.LoadLocation("Asia/Shanghai")
var DefaultTimeZone = defaultTimeZone
var defaultTimeFormat = "2006/01/02 15:04:05"
var defaultTimeMinFormat = "2006/01/02 15:04"
var backTimeFormat = "2006-01-02 15:04:05"

const defaultTimeStr = "1970/01/01 00:00:01"
const standTimeStr = "1970-01-01 00:00:01"

func DefaultFormatTime(t time.Time) string {
	if t.Unix() == 0 {
		return defaultTimeStr
	}
	return t.In(defaultTimeZone).Format(defaultTimeFormat)
}

func StandFormatTime(t time.Time) string {
	if t.Unix() == 0 {
		return standTimeStr
	}
	return t.In(defaultTimeZone).Format(backTimeFormat)
}

func DefaultFormatToTime(timeStr string) time.Time {
	t, err := time.ParseInLocation(defaultTimeFormat, timeStr, defaultTimeZone)
	if err != nil {
		t, err = time.ParseInLocation(backTimeFormat, timeStr, defaultTimeZone)
		if err != nil {
			return time.Now()
		}
	}
	return t
}

func DefaultFormatTimeMin(t time.Time) string {
	return t.In(defaultTimeZone).Format(defaultTimeMinFormat)
}

const defaultDateStr = "1970/01/01"

var defaultDateFormat = "2006/01/02"

func DefaultFormatDateNoSplit(t time.Time) string {
	if t.Unix() == 0 {
		return defaultDateStr
	}
	return t.In(defaultTimeZone).Format("20060102")
}

func DefaultFormatDate(t time.Time) string {
	if t.Unix() == 0 {
		return defaultDateStr
	}
	return t.In(defaultTimeZone).Format(defaultDateFormat)
}

func DefaultFormatDateStrToTime(timeStr string) time.Time {
	t, err := time.ParseInLocation(defaultDateFormat, timeStr, defaultTimeZone)
	if err != nil {
		return time.Now()
	}
	return t
}
