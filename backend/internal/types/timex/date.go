package timex

import (
	"database/sql/driver"
	"fmt"
	"github.com/spf13/cast"
	"strings"
	"time"
)

// DateStrTime 存储的是字符串但是又需要用于计算
type DateStrTime struct {
	timeStr   string // 字符换
	timestamp int64
}

func (d DateStrTime) GetTimeStamp() int64 {
	if d.timestamp > 0 {
		return d.timestamp
	}
	if d.timeStr != "" {
		return FormatDateStrToTime(d.timeStr).Unix()
	}
	return 0
}

func (d DateStrTime) GetDate() string {
	if d.timeStr == "" && d.timestamp == 0 {
		return ""
	}
	if d.timestamp > 0 {
		n := time.Now().Unix() * 10
		if d.timestamp > n {
			return FormatDate(time.Unix(d.timestamp/int64(time.Second), d.timestamp%int64(time.Second)))
		}
		return FormatDate(time.Unix(d.timestamp, 0))
	}
	if d.timeStr != "" {
		if len(d.timeStr) > 10 {
			return d.timeStr[:10]
		}
		return d.timeStr
	}
	return ""
}

func (d DateStrTime) Value() (driver.Value, error) {
	return d.GetDate(), nil
}

func (d *DateStrTime) Scan(v interface{}) error {
	switch v.(type) {
	case int64, int, int32, uint32, uint64:
		n := time.Now().Unix() * 10
		ux := cast.ToInt64(v)
		var vt time.Time
		if ux > n {
			vt = time.Unix(ux/int64(time.Second), ux%int64(time.Second))
		} else {
			vt = time.Unix(ux, 0)
		}
		*d = DateStrTime{
			timestamp: vt.Unix(),
			timeStr:   FormatDate(vt),
		}
		return nil
	case time.Time:
		v1 := v.(time.Time)
		*d = DateStrTime{
			timestamp: v1.Unix(),
			timeStr:   FormatDate(v1),
		}
		return nil
	case string, []byte:
		value := cast.ToString(v)
		if strings.Contains(value, "-") {
			v1 := FormatDateStrToTime(value)
			*d = DateStrTime{
				timestamp: v1.Unix(),
				timeStr:   FormatDate(v1),
			}
		} else {
			v1 := time.Time{}
			switch len(value) {
			case 10:
				v1 = time.Unix(cast.ToInt64(value), 0)
			case 16:
				v1 = time.Unix(cast.ToInt64(value)/int64(time.Second), 0)
			}

			*d = DateStrTime{
				timestamp: v1.Unix(),
				timeStr:   FormatDate(v1),
			}
		}
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (d DateStrTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", d.GetDate())), nil
}

func (d *DateStrTime) UnmarshalJSON(data []byte) error {
	if len(data) < 3 {
		return nil
	}
	if data[0] == '"' {
		data = data[1 : len(data)-1]
	}
	t := FormatDateStrToTime(string(data))
	*d = DateStrTime{
		timeStr:   FormatDate(t),
		timestamp: t.Unix(),
	}
	return nil
}
