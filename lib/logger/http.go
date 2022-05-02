package logger

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/utils"
	"github.com/rs/zerolog"
)

// https://learning-cloud-native-go.github.io/docs/a6.adding_zerolog_logger/

type httpLogger struct {
	l *zerolog.Logger
}

func NewHttpLogger() *httpLogger {
	l := NewDefaultLogger()
	return &httpLogger{
		l: l,
	}
}

func (hl *httpLogger) ReadBody(body io.ReadCloser) ([]byte, io.ReadCloser, error) {
	bodyBuffer, err := ioutil.ReadAll(body)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	reader := ioutil.NopCloser(bytes.NewBuffer(bodyBuffer))

	return bodyBuffer, reader, nil
}

func (hl *httpLogger) LogEntry(r *http.Request, start time.Time, body string) {
	hl.l.Info().
		Str("RequestId", utils.GenerateUUIDString()).
		Dur("Laytancy", time.Since(start)).
		Str("Protocol", r.Proto).
		Str("RequestMethod", r.Method).
		Str("Path", r.URL.Path).
		Str("RequestURL", r.URL.String()).
		Str("Query", r.URL.RawQuery).
		Str("Body", body).
		Str("Cookies", fmt.Sprint(r.Cookies())).
		Str("Origin", r.Header.Get("Origin")).
		Str("UserAgent", r.UserAgent()).
		Str("Referer", r.Referer()).
		Str("ClientIp", clientIP(r.Header)).
		Send()
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
