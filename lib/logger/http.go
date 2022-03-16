package logger

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/utils"
	"github.com/rs/zerolog"
)

// https://learning-cloud-native-go.github.io/docs/a6.adding_zerolog_logger/

type HttpLogger interface {
	LogEntry(r *http.Request, start time.Time)
}

type httpLogger struct {
	z *zerolog.Logger
}

func NewHttpLogger() HttpLogger {
	z := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &httpLogger{
		z: &z,
	}
}

func (hl *httpLogger) LogEntry(r *http.Request, start time.Time) {
	bodyBuf, reader, err := readBody(r.Body)
	if err != nil {
		Logger.
			Err(err).
			Str("Action", "RequestLog Read Body").
			Msgf("Params: %+v", r.Body)
	}
	r.Body = reader

	hl.z.Info().
		Dur("Laytancy", time.Since(start)).
		Str("Protocol", r.Proto).
		Str("RequestId", utils.GenerateUUIDString()).
		Str("RequestMethod", r.Method).
		Str("Path", r.URL.Path).
		Str("RequestURL", r.URL.String()).
		Str("Query", r.URL.RawQuery).
		Str("Body", string(bodyBuf)).
		Str("Cookies", fmt.Sprint(r.Cookies())).
		Str("Origin", r.Header.Get("Origin")).
		Str("UserAgent", r.UserAgent()).
		Str("Referer", r.Referer()).
		Str("ClientIp", clientIP(r.Header)).
		Msg("")
}

var (
	trueClientIP          = http.CanonicalHeaderKey("True-Client-IP")
	xForwardedFor         = http.CanonicalHeaderKey("X-Forwarded-For")
	xRealIP               = http.CanonicalHeaderKey("X-Real-IP")
	xEnvoyExternalAddress = http.CanonicalHeaderKey("X-Envoy-External-Address")
)

func readBody(body io.ReadCloser) ([]byte, io.ReadCloser, error) {
	bodyBuf, err := ioutil.ReadAll(body)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	reader := ioutil.NopCloser(bytes.NewBuffer(bodyBuf))

	return bodyBuf, reader, nil
}

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
