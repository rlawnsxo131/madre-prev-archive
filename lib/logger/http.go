package logger

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// go get -u github.com/rs/zerolog/log

type HttpLogger interface {
	RequestLog(r *http.Request)
	ResponseLog(status int, data interface{})
}

type httpLogger struct {
	z *zerolog.Logger
}

type logEntry struct {
	ReceivedTime       time.Time
	RequestMethod      string
	RequestURL         string
	RequestHeaderSize  int64
	RequestBodySize    int64
	UserAgent          string
	Referer            string
	Proto              string
	RemoteIP           string
	ServerIP           string
	Status             int
	ResponseHeaderSize int64
	ResponseBodySize   int64
	Latency            time.Duration
}

func NewHttpLogger() HttpLogger {
	z := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &httpLogger{
		z: &z,
	}
}

func (hl *httpLogger) RequestLog(r *http.Request) {
	// start := time.Now()
	// le := &logEntry{
	// 	ReceivedTime:      start,
	//     RequestMethod:     r.Method,
	//     RequestURL:        r.URL.String(),
	//     RequestHeaderSize: headerSize(r.Header),
	//     UserAgent:         r.UserAgent(),
	//     Referer:           r.Referer(),
	//     Proto:             r.Proto,
	//     RemoteIP:          ipFromHostPort(r.RemoteAddr),
	// }
	clientIp := clientIP(r.Header)
	hl.z.Info().
		Str("Protocol", r.Proto).
		Str("Origin", r.Header.Get("Origin")).
		Str("Method", r.Method).
		Str("Path", r.URL.Path).
		Str("ClientIp", clientIp).
		Msg("")
}

func (hl *httpLogger) ResponseLog(status int, data interface{}) {}

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
