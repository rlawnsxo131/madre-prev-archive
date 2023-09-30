package logger

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
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

// thread safe singleton default logger
var (
	_onceDefaultLogger sync.Once
	defaultLogger      *zerolog.Logger
)

func DefaultLogger() *zerolog.Logger {
	_onceDefaultLogger.Do(func() {
		wr := diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
			fmt.Printf("Logger Dropped %d messages", missed)
		})
		l := zerolog.New(wr)
		defaultLogger = &l
	})
	return defaultLogger
}

// new instance logger
func NewDefaultLogger() *zerolog.Logger {
	l := zerolog.New(os.Stdout)
	return &l
}
