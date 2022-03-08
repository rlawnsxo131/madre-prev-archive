package logger

import (
	"net/http"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// go get -u github.com/rs/zerolog/log

type HttpLogger interface{}

type httpLogger struct {
	z *zerolog.Logger
}

func NewHttpLogger() *httpLogger {
	z := zerolog.New(os.Stderr)
	return &httpLogger{
		z: &z,
	}
}

func (l *httpLogger) Logging(r *http.Request) {
	clientIp := clientIP(r.Header)
	l.z.Info().
		Str("clientIp", clientIp).
		Msg("")
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
