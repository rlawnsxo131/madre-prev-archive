package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
