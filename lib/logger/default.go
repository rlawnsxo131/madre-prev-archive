package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func NewDefaultLogger() *zerolog.Logger {
	l := zerolog.New(os.Stderr).With().Timestamp().Logger() //.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	return &l
}
