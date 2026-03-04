package logger

import (
	"context"
	"log/slog"
	"os"
)

type ctxKey string
const TraceIDKey ctxKey = "trace_id"

func New(level string) *slog.Logger {
	opts := &slog.HandlerOptions{Level: parseLevel(level)}
	return slog.New(slog.NewJSONHandler(os.Stdout, opts))
}

func parseLevel(level string) slog.Level {
	switch level {
	case "debug": return slog.LevelDebug
	case "warn":  return slog.LevelWarn
	case "error": return slog.LevelError
	default:      return slog.LevelInfo
	}
}

func WithCtx(ctx context.Context, l *slog.Logger) *slog.Logger {
	if tid, ok := ctx.Value(TraceIDKey).(string); ok {
		return l.With("trace_id", tid)
	}
	return l
}
