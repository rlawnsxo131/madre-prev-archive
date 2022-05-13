package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func NewBaseLogger() *zerolog.Logger {
	//.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	l := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &l
}
