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

func (hle *httpLogEntry) ReadBody() error {
	if hle.r.Body != nil {
		body, err := io.ReadAll(hle.r.Body)
		if err != nil {
			hle.add = append(hle.add, func(e *zerolog.Event) {
				e.Err(errors.Wrap(err, "read http body error"))
			})
			return err
		}
		hle.body = append(hle.body, body...)
		hle.r.Body = io.NopCloser(
			bytes.NewBuffer(body),
		)
	}
	return nil
}

func (hle *httpLogEntry) Add(f func(e *zerolog.Event)) {
	hle.add = append(hle.add, f)
}

func (hle *httpLogEntry) Write(t time.Time) {
	e := hle.l.Log().Timestamp().
		Str("requestId", chi_middleware.GetReqID(hle.r.Context())).
		Dur("elapsed(ms)", time.Since(t)).
		Str("protocol", hle.r.Proto).
		Str("method", hle.r.Method).
		Str("uri", hle.r.URL.RequestURI()).
		Bytes("body", hle.body).
		Str("origin", hle.r.Header.Get("Origin")).
		Str("referer", hle.r.Referer()).
		Int("status", hle.ww.Status()).
		Str("agent", hle.r.UserAgent()).
		Str("remoteAddr", hle.r.RemoteAddr).
		Str("cookies", fmt.Sprint(hle.r.Cookies()))

	for _, f := range hle.add {
		f(e)
	}

	e.Send()
}

type key int

const (
	KEY_LOG_ENTRY_CTX key = iota
)

func LogEntry(ctx context.Context) HTTPLogEntry {
	v := ctx.Value(KEY_LOG_ENTRY_CTX)
	if v, ok := v.(HTTPLogEntry); ok {
		return v
	}
	return nil
}

func SetLogEntry(ctx context.Context, le HTTPLogEntry) context.Context {
	return context.WithValue(
		ctx,
		KEY_LOG_ENTRY_CTX,
		le,
	)
}
