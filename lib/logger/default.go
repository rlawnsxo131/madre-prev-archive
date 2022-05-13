package logger

import (
	"sync"

	"github.com/rs/zerolog"
)

var (
	defaultLogger     *zerolog.Logger
	onceDefaultLogger sync.Once
)

func GetDefaultLogger() *zerolog.Logger {
	onceDefaultLogger.Do(func() {
		defaultLogger = NewBaseLogger()
	})
	return defaultLogger
}
