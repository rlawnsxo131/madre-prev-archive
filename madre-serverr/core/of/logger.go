package of

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

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
