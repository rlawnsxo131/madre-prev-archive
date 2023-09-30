package logger

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rlawnsxo131/madre-server/core/errorz"
	"github.com/rs/zerolog"
)

var DefaultHTTPLogger = NewHTTPLogger(os.Stdout, NewHTTPLogFormatter())

// http logger
type HTTPLogger struct {
	l *zerolog.Logger
	f *httpLogFormatter
}

func NewHTTPLogger(w io.Writer, f *httpLogFormatter) *HTTPLogger {
	l := zerolog.New(w)
	return &HTTPLogger{
		l: &l,
		f: f,
	}
}

func (hl *HTTPLogger) NewLogEntry(r *http.Request, ww chi_middleware.WrapResponseWriter) LogEntry {
	return hl.f.NewLogEntry(hl.l, r, ww)
}

// formatter
type httpLogFormatter struct{}

func NewHTTPLogFormatter() *httpLogFormatter {
	return &httpLogFormatter{}
}

func (hlf *httpLogFormatter) NewLogEntry(
	l *zerolog.Logger,
	r *http.Request,
	ww chi_middleware.WrapResponseWriter,
) LogEntry {
	return &httpLogEntry{
		l:    l,
		r:    r,
		ww:   ww,
		body: []byte{},
		add:  []func(e *zerolog.Event){},
	}
}

// entry
type httpLogEntry struct {
	l    *zerolog.Logger
	r    *http.Request
	ww   chi_middleware.WrapResponseWriter
	body []byte
	add  []func(e *zerolog.Event)
}

func (le *httpLogEntry) ReadBody() error {
	if le.r.Body != nil {
		body, err := io.ReadAll(le.r.Body)
		if err != nil {
			le.add = append(le.add, func(e *zerolog.Event) {
				e.Err(
					errorz.New(
						fmt.Errorf("readAll error: %+v", err),
					),
				)
			})
			return err
		}
		le.body = append(le.body, body...)
		le.r.Body = io.NopCloser(
			bytes.NewBuffer(body),
		)
	}
	return nil
}

func (le *httpLogEntry) Add(f func(e *zerolog.Event)) {
	le.add = append(le.add, f)
}

func (le *httpLogEntry) Write(t time.Time) {
	var (
		e          *zerolog.Event
		statusCode int = le.ww.Status()
	)

	switch {
	case statusCode > 0 && statusCode < 300:
		e = le.l.Info()
	case statusCode > 299 && statusCode < 500:
		e = le.l.Warn()
	case statusCode > 499:
		e = le.l.Error()
	default:
		e = le.l.Error()
	}

	e.Str("time", t.UTC().Format(time.RFC3339Nano)).
		Str("requestId", chi_middleware.GetReqID(le.r.Context())).
		Dur("elapsed(ms)", time.Since(t)).
		Str("protocol", le.r.Proto).
		Str("method", le.r.Method).
		Str("uri", le.r.URL.RequestURI()).
		Bytes("body", le.body).
		Str("origin", le.r.Header.Get("Origin")).
		Str("referer", le.r.Referer()).
		Int("status", le.ww.Status()).
		Str("cookies", fmt.Sprint(le.r.Cookies())).
		Str("userAgent", le.r.UserAgent()).
		Str("remoteAddr", le.r.RemoteAddr)

	for _, f := range le.add {
		f(e)
	}

	e.Send()
}
