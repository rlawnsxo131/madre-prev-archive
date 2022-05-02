package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type defaultLogger struct {
	Logger *zerolog.Logger
}

func NewDefaultLogger() *zerolog.Logger {
	l := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &l
}
