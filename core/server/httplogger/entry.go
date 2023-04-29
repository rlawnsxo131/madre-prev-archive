package httplogger

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type HTTPLogEntry interface {
	ReadBody() error
	Add(func(e *zerolog.Event))
	Write(t time.Time)
}

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
				e.Err(errors.Wrap(err, "read http body error"))
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
	e := le.l.Log().Timestamp().
		Str("requestId", chi_middleware.GetReqID(le.r.Context())).
		Dur("elapsed(ms)", time.Since(t)).
		Str("protocol", le.r.Proto).
		Str("method", le.r.Method).
		Str("uri", le.r.URL.RequestURI()).
		Bytes("body", le.body).
		Str("origin", le.r.Header.Get("Origin")).
		Str("referer", le.r.Referer()).
		Int("status", le.ww.Status()).
		Str("agent", le.r.UserAgent()).
		Str("remoteAddr", le.r.RemoteAddr).
		Str("cookies", fmt.Sprint(le.r.Cookies()))

	for _, f := range le.add {
		f(e)
	}

	e.Send()
}

type key int

const (
	KEY_LOG_ENTRY_CTX key = iota
)

func LogEntry(ctx context.Context) (HTTPLogEntry, error) {
	v := ctx.Value(KEY_LOG_ENTRY_CTX)
	if v, ok := v.(HTTPLogEntry); ok {
		return v, nil
	}
	return nil, errors.New("there is no httpLogEntry in the context")
}

func SetLogEntry(ctx context.Context, le HTTPLogEntry) context.Context {
	return context.WithValue(
		ctx,
		KEY_LOG_ENTRY_CTX,
		le,
	)
}
