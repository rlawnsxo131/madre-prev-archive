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

func DefaultLogger() *zerolog.Logger {
	onceDefaultLogger.Do(func() {
		l := zerolog.New(os.Stdout).With().Logger()
		defaultLogger = &l
	})
	return defaultLogger
}
