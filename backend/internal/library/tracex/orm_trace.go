package tracex

import (
	"context"
	"github.com/zeromicro/go-zero/core/trace"
	"go.opentelemetry.io/otel/attribute"
	otelTrace "go.opentelemetry.io/otel/trace"
	"time"
)

const (
	DBSpanNameKey    = "db_query"
	DBSqlKey         = "query"
	DBTimeKey        = "time"
	DBAffectedKey    = "affected"
	DBErrorKey       = "error"
	DBDescriptionKey = "description"
)

func OrmTrace(ctx context.Context, begin time.Time, sqlStr string, rowsAffected int64, err error) {
	tracer := trace.TracerFromContext(ctx)
	_, span := tracer.Start(ctx, DBSpanNameKey, otelTrace.WithTimestamp(begin))
	defer func() {
		span.End()
	}()
	span.SetAttributes(attribute.KeyValue{
		Key:   DBSqlKey,
		Value: attribute.StringValue(sqlStr),
	})
	span.SetAttributes(attribute.Int64(DBTimeKey, time.Now().Sub(begin).Milliseconds()))
	span.SetAttributes(attribute.Int64(DBAffectedKey, rowsAffected))
	if err != nil {
		span.SetAttributes(attribute.KeyValue{
			Key:   DBErrorKey,
			Value: attribute.BoolValue(true),
		})
		span.SetAttributes(attribute.KeyValue{
			Key:   DBDescriptionKey,
			Value: attribute.StringValue(err.Error()),
		})
	}
}
