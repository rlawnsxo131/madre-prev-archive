package logger

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type logEntryCtxKey int

var _key = logEntryCtxKey(1)

type LogEntry interface {
	ReadBody() error
	Add(func(e *zerolog.Event))
	Write(t time.Time)
}

func GetLogEntry(ctx context.Context) LogEntry {
	if le, ok := ctx.Value(_key).(LogEntry); ok {
		return le
	}
	return nil
}

func SetLogEntry(ctx context.Context, le LogEntry) context.Context {
	return context.WithValue(ctx, _key, le)
}

func NewDefaultZeroLogger() *zerolog.Logger {
	l := zerolog.New(os.Stdout)
	return &l
}
