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

func NewDefaultLogger() *zerolog.Logger {
	onceDefaultLogger.Do(func() {
		//.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		l := zerolog.New(os.Stderr).With().Timestamp().Logger()
		defaultLogger = &l
	})
	return defaultLogger
}
