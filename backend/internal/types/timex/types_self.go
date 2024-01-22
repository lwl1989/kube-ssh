package timex

import (
	"database/sql/driver"
	"fmt"
	"github.com/go-libraries/kube-manager/backend/internal/library/utils"
	"github.com/spf13/cast"
	"time"
)

type SelfTime time.Time

var defaultTimeFormat = "2006/01/02 15:04:05"
var backTimeFormat = "2006-01-02 15:04:05"
var ZeroTime = SelfTime(time.Unix(1, 0))
var EmptyStrByt = []byte{'"', '"'}

func init() {
	diffTime = utils.DefaultFormatToTime("2000/01/01 00:00:01")
}

func (t SelfTime) String() string {
	return utils.DefaultFormatTime(time.Time(t))
}

func (t SelfTime) ToDate() string {
	return utils.DefaultFormatDate(time.Time(t))
}

func (t SelfTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Unix(t.Unix(), 0).In(utils.DefaultTimeZone)
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return zeroTime, nil
	}
	return tlt, nil
}

func (t *SelfTime) Scan(v interface{}) error {
	switch v.(type) {
	case int64, int, int32, uint32, uint64:
		*t = SelfTime(time.Unix(cast.ToInt64(v), 0))
		return nil
	case time.Time:
		*t = SelfTime(v.(time.Time).In(utils.DefaultTimeZone))
		return nil
	case string:
		*t = SelfTime(StringToTime(v.(string)))
	case []byte:
		value := v.([]byte)
		var t1 time.Time
		if len(value) == 10 {
			t1 = time.Unix(cast.ToInt64(value), 0)
			*t = SelfTime(t1)
		} else if len(value) == 16 {
			t1 = time.Unix(cast.ToInt64(value), 0)
			*t = SelfTime(t1)
		} else {
			*t = SelfTime(StringToTime(string(value)))
		}
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t SelfTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(t)
	if tTime.IsZero() || tTime.Unix() == 0 {
		return EmptyStrByt, nil
	}
	if t.Unix() < diffTime.Unix() {
		return EmptyStrByt, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", utils.DefaultFormatTime(tTime))), nil
}

func (t *SelfTime) UnmarshalJSON(data []byte) error {
	if len(data) < 3 {
		return nil
	}
	if data[0] == '"' {
		data = data[1 : len(data)-1]
	}

	*t = SelfTime(utils.DefaultFormatToTime(string(data)))
	return nil
}

func TimeNow() SelfTime {
	return SelfTime(time.Now())
}

func (t SelfTime) Unix() int64 {
	return time.Time(t).Unix()
}

func StringToTime(timeStr string) time.Time {
	t, err := time.ParseInLocation(defaultTimeFormat, timeStr, utils.DefaultTimeZone)
	if err != nil {
		t, err = time.ParseInLocation(backTimeFormat, timeStr, utils.DefaultTimeZone)
		if err != nil {
			return time.Now()
		}
	}
	return t
}
