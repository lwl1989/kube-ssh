package timex

import (
	"database/sql/driver"
	"fmt"
	library "github.com/lwl1989/kube-ssh/backend/internal/library/utils"
	"github.com/spf13/cast"
	"time"
)

var diffTime time.Time

type UnixTime int64

func (t UnixTime) String() string {
	return library.DefaultFormatTime(t.ToTime())
}

func (t UnixTime) Value() (driver.Value, error) {
	return int64(t), nil
}

func (t *UnixTime) Scan(v interface{}) error {
	switch v.(type) {
	case int64, int, int32, uint32, uint64:
		*t = UnixTime(cast.ToInt64(v))
		return nil
	case []byte:
		value := v.([]byte)
		var t1 time.Time
		if len(value) == 10 {
			t1 = time.Unix(cast.ToInt64(value), 0)
		}
		if len(value) == 16 {
			t1 = time.Unix(cast.ToInt64(value), 0)
		}
		*t = UnixTime(t1.Unix())
	case string:
		value := v.(string)
		t1 := cast.ToInt64(value)
		if t1 > 0 {
			*t = UnixTime(t1)
		} else {
			return fmt.Errorf("can not convert %v to timestamp", v)
		}
	default:
		return fmt.Errorf("can not convert %v to timestamp", v)
	}
	return nil
}

func (t UnixTime) MarshalJSON() ([]byte, error) {
	tTime := time.Unix(int64(t), 0)
	if tTime.IsZero() || tTime.Unix() == 0 {
		return EmptyStrByt, nil
	}
	if t.Unix() < diffTime.Unix() {
		return EmptyStrByt, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", library.DefaultFormatTime(tTime))), nil
}

func (t *UnixTime) UnmarshalJSON(data []byte) error {
	if len(data) < 3 {
		return nil
	}
	if data[0] == '"' {
		data = data[1 : len(data)-1]
	}

	*t = UnixTime(cast.ToInt(library.DefaultFormatToTime(string(data)).Unix()))
	return nil
}

func UnixTimeNow() UnixTime {
	return UnixTime(int(time.Now().Unix()))
}

func (t UnixTime) Unix() int64 {
	return int64(t)
}

func (t UnixTime) ToTime() time.Time {
	return time.Unix(t.Unix(), 0)
}

func (t UnixTime) ToDate() string {
	return library.DefaultFormatDate(t.ToTime())
}
