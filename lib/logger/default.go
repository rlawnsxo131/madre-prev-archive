package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type defaultLogger struct {
	l   *zerolog.Logger
	add []func(e *zerolog.Event)
}

func NewDefaultLogger() *defaultLogger {
	l := zerolog.New(os.Stdout).With().Logger()
	return &defaultLogger{
		l:   &l,
		add: []func(e *zerolog.Event){},
	}
}

func (dl *defaultLogger) Add(f func(e *zerolog.Event)) *defaultLogger {
	dl.add = append(dl.add, f)
	return dl
}

func (dl *defaultLogger) Send() {
	e := dl.l.Log().Timestamp()
	for _, f := range dl.add {
		f(e)
	}
	e.Send()
}

func (dl *defaultLogger) SendInfo() {
	e := dl.l.Info().Timestamp()
	for _, f := range dl.add {
		f(e)
	}
	e.Send()
}

func (dl *defaultLogger) SendError() {
	e := dl.l.Error().Timestamp()
	for _, f := range dl.add {
		f(e)
	}
	e.Send()
}

func (dl *defaultLogger) SendFatal() {
	e := dl.l.Fatal().Timestamp()
	for _, f := range dl.add {
		f(e)
	}
	e.Send()
}
