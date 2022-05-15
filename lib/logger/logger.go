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

func NewBaseLogger() *zerolog.Logger {
	//.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	l := zerolog.New(os.Stderr).With().Logger()
	return &l
}

func DefaultLogger() *zerolog.Logger {
	onceDefaultLogger.Do(func() {
		defaultLogger = NewBaseLogger()
	})
	return defaultLogger
}
