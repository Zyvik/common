package logging

import (
	"context"

	commonCtx "github.com/Zyvik/common/context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	LoggerCtxKey  = commonCtx.Key("logger")
	TraceIDCtxKey = commonCtx.Key("traceID")
	SpanIDCtxKey  = commonCtx.Key("spanID")
)

func AddLoggerToContext(ctx context.Context, log *logrus.Logger) context.Context {
	return context.WithValue(ctx, LoggerCtxKey, log)
}

func AddTracingToContext(ctx context.Context, traceID, spanID string) context.Context {
	if traceID == "" {
		traceID = uuid.NewString()
	}
	if spanID == "" {
		spanID = uuid.NewString()
	}
	return context.WithValue(context.WithValue(ctx, TraceIDCtxKey, traceID), SpanIDCtxKey, spanID)
}

func GetTracedEntry(ctx context.Context) *logrus.Entry {
	log, ok := ctx.Value(LoggerCtxKey).(*logrus.Logger)
	if !ok {
		log = logrus.New()
	}

	return log.WithFields(map[string]interface{}{
		TraceID: commonCtx.GetString(ctx, TraceIDCtxKey),
		SpanID:  commonCtx.GetString(ctx, SpanIDCtxKey),
	})
}
