package httplogger

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type httpLogger struct {
	l    *zerolog.Logger
	r    *http.Request
	ww   chi_middleware.WrapResponseWriter
	body []byte
	add  []func(e *zerolog.Event)
}

func NewHTTPLogger(r *http.Request, ww chi_middleware.WrapResponseWriter) *httpLogger {
	l := zerolog.New(os.Stdout).With().Logger()
	return &httpLogger{
		l:    &l,
		r:    r,
		ww:   ww,
		body: []byte{},
		add:  []func(e *zerolog.Event){},
	}
}

func (hl *httpLogger) ReadBody() error {
	if hl.r.Body != nil {
		body, err := io.ReadAll(hl.r.Body)
		if err != nil {
			hl.add = append(hl.add, func(e *zerolog.Event) {
				e.Err(errors.Wrap(err, "read http body error"))
			})
			return err
		}
		hl.body = append(hl.body, body...)
		hl.r.Body = io.NopCloser(
			bytes.NewBuffer(body),
		)
	}
	return nil
}

func (hl *httpLogger) Add(f func(e *zerolog.Event)) {
	hl.add = append(hl.add, f)
}

func (hl *httpLogger) Write(t time.Time) {
	e := hl.l.Log().Timestamp().
		Str("requestId", chi_middleware.GetReqID(hl.r.Context())).
		Dur("elapsed(ms)", time.Since(t)).
		Str("protocol", hl.r.Proto).
		Str("method", hl.r.Method).
		Str("uri", hl.r.URL.RequestURI()).
		Bytes("body", hl.body).
		Str("origin", hl.r.Header.Get("Origin")).
		Str("referer", hl.r.Referer()).
		Int("status", hl.ww.Status()).
		Str("agent", hl.r.UserAgent()).
		Str("remoteAddr", hl.r.RemoteAddr).
		Str("cookies", fmt.Sprint(hl.r.Cookies()))

	for _, f := range hl.add {
		f(e)
	}

	e.Send()
}

type key int

const (
	KEY_LOGGER_CTX key = iota
)

func HTTPLoggerCtx(ctx context.Context) *httpLogger {
	v := ctx.Value(KEY_LOGGER_CTX)
	if v, ok := v.(*httpLogger); ok {
		return v
	}
	return nil
}

func SetHTTPLoggerCtx(ctx context.Context, hl *httpLogger) context.Context {
	return context.WithValue(
		ctx,
		KEY_LOGGER_CTX,
		hl,
	)
}
