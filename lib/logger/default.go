package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
)

var (
	defaultLogger     *zerolog.Logger
	onceDefaultLogger sync.Once
)

func GetDefaultLogger() *zerolog.Logger {
	onceDefaultLogger.Do(func() {
		l := NewDefaultLogger()
		defaultLogger = l
	})
	return defaultLogger
}

func NewDefaultLogger() *zerolog.Logger {
	//.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	l := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &l
}
