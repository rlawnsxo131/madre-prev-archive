package httplogger

import (
	"io"
	"net/http"
	"os"

	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

var DefaultHTTPLogger = NewHTTPLogger(os.Stdout, NewHTTPLogFormatter())

type HTTPLogger struct {
	l *zerolog.Logger
	f HTTPLogFormatter
}

func NewHTTPLogger(w io.Writer, f HTTPLogFormatter) *HTTPLogger {
	l := zerolog.New(w)
	return &HTTPLogger{
		l: &l,
		f: f,
	}
}

func (hl *HTTPLogger) NewLogEntry(r *http.Request, ww chi_middleware.WrapResponseWriter) HTTPLogEntry {
	return hl.f.NewLogEntry(hl.l, r, ww)
}

// formatter
type HTTPLogFormatter interface {
	NewLogEntry(l *zerolog.Logger, r *http.Request, ww chi_middleware.WrapResponseWriter) HTTPLogEntry
}

type httpLogFormatter struct{}

func NewHTTPLogFormatter() HTTPLogFormatter {
	return &httpLogFormatter{}
}

func (hlf *httpLogFormatter) NewLogEntry(l *zerolog.Logger, r *http.Request, ww chi_middleware.WrapResponseWriter) HTTPLogEntry {
	return &httpLogEntry{
		l:    l,
		r:    r,
		ww:   ww,
		body: []byte{},
		add:  []func(e *zerolog.Event){},
	}
}
