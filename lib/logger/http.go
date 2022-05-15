package logger

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type HTTPLogger interface {
	Add(f func(e *zerolog.Event))
	Write(t time.Time)
}

type httpLogger struct {
	l   *zerolog.Logger
	r   *http.Request
	add []func(e *zerolog.Event)
}

func NewHTTPLogger(r *http.Request) HTTPLogger {
	return &httpLogger{
		l:   NewBaseLogger(),
		r:   r,
		add: []func(e *zerolog.Event){},
	}
}

func (hl *httpLogger) Add(f func(e *zerolog.Event)) {
	hl.add = append(hl.add, f)
}

func (hl *httpLogger) Write(t time.Time) {
	hl.r.Cookies()
	e := hl.l.Log().Timestamp().
		Str("requestId", uuid.NewString()).
		Dur("elapsed(ms)", time.Since(t)).
		Str("protocol", hl.r.Proto).
		Str("method", hl.r.Method).
		Str("origin", hl.r.Header.Get("Origin")).
		Str("uri", hl.r.URL.RequestURI()).
		Str("agent", hl.r.UserAgent()).
		Str("referer", hl.r.Referer()).
		Str("clientIp", clientIP(hl.r.Header))

	for _, f := range hl.add {
		f(e)
	}

	e.Str("cookies", fmt.Sprint(hl.r.Cookies()))
	e.Msg("")
}

var (
	trueClientIP          = http.CanonicalHeaderKey("True-Client-IP")
	xForwardedFor         = http.CanonicalHeaderKey("X-Forwarded-For")
	xRealIP               = http.CanonicalHeaderKey("X-Real-IP")
	xEnvoyExternalAddress = http.CanonicalHeaderKey("X-Envoy-External-Address")
)

// clientIP returns the IP of the client.
// If a header identifying the real IP exists, the value of the header will be used.
func clientIP(h http.Header) string {
	if tcip := h.Get(trueClientIP); tcip != "" {
		return tcip
	} else if xrip := h.Get(xRealIP); xrip != "" {
		return xrip
	} else if xff := h.Get(xForwardedFor); xff != "" {
		i := strings.Index(xff, ",")
		if i == -1 {
			i = len(xff)
		}
		return xff[:i]
	} else if xeea := h.Get(xEnvoyExternalAddress); xeea != "" {
		return xeea
	}

	return ""
}

const (
	Key_HTTPLoggerCtx = "Key_HTTPLoggerCtx"
)

func HTTPLoggerCtx(ctx context.Context) HTTPLogger {
	v := ctx.Value(Key_HTTPLoggerCtx)
	if v, ok := v.(HTTPLogger); ok {
		return v
	}
	return nil
}

func SetHTTPLoggerCtx(ctx context.Context, hl HTTPLogger) context.Context {
	return context.WithValue(ctx, Key_HTTPLoggerCtx, hl)
}
