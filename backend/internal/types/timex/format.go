package timex

import (
	"time"
)

var defaultTimeZone, _ = time.LoadLocation("Asia/Shanghai")

var defaultTimeMinFormat = "2006/01/02 15:04"
var defaultDateFormat = "2006/01/02"

const defaultTimeStr = "1970/01/01 00:00:01"
const defaultDateStr = "1970/01/01"

func FormatTimeToString(t time.Time) string {
	if t.Unix() == 0 {
		return defaultTimeStr
	}
	return t.In(defaultTimeZone).Format(defaultTimeFormat)
}

func FormatTimeMin(t time.Time) string {
	return t.In(defaultTimeZone).Format(defaultTimeMinFormat)
}

func FormatDate(t time.Time) string {
	if t.Unix() == 0 {
		return defaultDateStr
	}
	return t.In(defaultTimeZone).Format(defaultDateFormat)
}

func FormatDateStrToTime(timeStr string) time.Time {
	if len(timeStr) > 10 {
		return StringToTime(timeStr)
	}
	t, err := time.ParseInLocation(defaultDateFormat, timeStr, defaultTimeZone)
	if err != nil {
		return time.Now()
	}
	return t
}
