package logger

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/utils"
	"github.com/rs/zerolog"
)

type HTTPLogger interface {
	ReadBody(r *http.Request) ([]byte, error)
	Add(f func(e *zerolog.Event))
	Write(t time.Time, body string)
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

func (hl *httpLogger) ReadBody(r *http.Request) ([]byte, error) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	reader := ioutil.NopCloser(
		bytes.NewBuffer(buf),
	)
	r.Body = reader
	return buf, nil
}

func (hl *httpLogger) Add(f func(e *zerolog.Event)) {
	hl.add = append(hl.add, f)
}

func (hl *httpLogger) Write(t time.Time, body string) {
	e := hl.l.Log().
		Str("RequestId", utils.GenerateUUIDString()).
		Dur("Laytancy", time.Since(t)).
		Str("Protocol", hl.r.Proto).
		Str("RequestMethod", hl.r.Method).
		Str("Path", hl.r.URL.Path).
		Str("Query", hl.r.URL.RawQuery).
		Str("Body", body).
		Str("Cookies", fmt.Sprint(hl.r.Cookies())).
		Str("Origin", hl.r.Header.Get("Origin")).
		Str("UserAgent", hl.r.UserAgent()).
		Str("Referer", hl.r.Referer()).
		Str("ClientIp", clientIP(hl.r.Header))

	for _, f := range hl.add {
		f(e)
	}

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
