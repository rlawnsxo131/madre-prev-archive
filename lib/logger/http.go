package logger

import (
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type HTTPLogger struct {
	l   *zerolog.Logger
	r   *http.Request
	add []func(e *zerolog.Event)
}

func NewHTTPLogger(r *http.Request) *HTTPLogger {
	return &HTTPLogger{
		l:   NewBaseLogger(),
		r:   r,
		add: []func(e *zerolog.Event){},
	}
}

func (hl *HTTPLogger) Add(f func(e *zerolog.Event)) {
	hl.add = append(hl.add, f)
}

func (hl *HTTPLogger) Write(t time.Time) {
	// e := hl.l.Log().
	// 	Str("protocol", "http").
	// 	Str("path", hl.r.URL.Path).
	// 	Str("time", t.UTC().Format(time.RFC3339Nano)).
	// 	Dur("elapsed(ms)", time.Since(t))

	// e := hl.l.Log().
	// 	Str("RequestId", reqId).
	// 	Dur("Laytancy", time.Since(start)).
	// 	Str("Protocol", r.Proto).
	// 	Str("RequestMethod", r.Method).
	// 	Str("Path", r.URL.Path).
	// 	Str("RequestURL", r.URL.String()).
	// 	Str("Query", r.URL.RawQuery).
	// 	Str("Body", body).
	// 	Str("Cookies", fmt.Sprint(r.Cookies())).
	// 	Str("Origin", r.Header.Get("Origin")).
	// 	Str("UserAgent", r.UserAgent()).
	// 	Str("Referer", r.Referer()).
	// 	Str("ClientIp", clientIP(r.Header)).

	// for _, f := range le.add {
	// 	f(e)
	// }

	// e.Send()
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
