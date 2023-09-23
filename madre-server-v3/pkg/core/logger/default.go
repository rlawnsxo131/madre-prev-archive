package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var DefaultLogger = NewDefaultLogger(os.Stdout)

type defaultLogger struct {
	l *zerolog.Logger
}

func NewDefaultLogger(w io.Writer) *defaultLogger {
	l := zerolog.New(w)
	return &defaultLogger{
		l: &l,
	}
}

func (dl *defaultLogger) NewLogEntry() *defaultLogEntry {
	return &defaultLogEntry{
		l:   dl.l,
		add: []func(e *zerolog.Event){},
	}
}

type defaultLogEntry struct {
	l   *zerolog.Logger
	add []func(e *zerolog.Event)
}

func (dle *defaultLogEntry) Add(f func(e *zerolog.Event)) *defaultLogEntry {
	dle.add = append(dle.add, f)
	return dle
}

func (dle *defaultLogEntry) Send() {
	e := dle.l.Log().Str("time", time.Now().UTC().Format(time.RFC3339Nano))
	for _, f := range dle.add {
		f(e)
	}
	e.Send()
}

func (dle *defaultLogEntry) SendInfo() {
	e := dle.l.Info().Str("time", time.Now().UTC().Format(time.RFC3339Nano))
	for _, f := range dle.add {
		f(e)
	}
	e.Send()
}

func (dle *defaultLogEntry) SendError() {
	e := dle.l.Error().Str("time", time.Now().UTC().Format(time.RFC3339Nano))
	for _, f := range dle.add {
		f(e)
	}
	e.Send()
}

func (dle *defaultLogEntry) SendFatal() {
	e := dle.l.Fatal().Str("time", time.Now().UTC().Format(time.RFC3339Nano))
	for _, f := range dle.add {
		f(e)
	}
	e.Send()
}
