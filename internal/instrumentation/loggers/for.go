package loggers

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
)

func For(ctx context.Context) *zap.Logger {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		if jaegerCtx, ok := span.Context().(jaeger.SpanContext); ok {
			return loggers.With([]zap.Field{
				zap.String("trace_id", jaegerCtx.TraceID().String()),
				zap.String("span_id", jaegerCtx.SpanID().String()),
			}...)
		}
	}
	return loggers
}
