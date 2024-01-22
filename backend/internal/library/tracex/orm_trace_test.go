package tracex

import (
	"context"
	"github.com/go-libraries/ormx"
	"testing"
	"time"
)

func TestOrmTrace(t *testing.T) {
	lg := ormx.NewLog("info")
	lg.SetTrace(OrmTrace)
	lg.Trace(context.Background(), time.Now().Add(time.Millisecond*-100), func() (sql string, rowsAffected int64) {
		return "select * from users limit 10", 10
	}, nil)
}
