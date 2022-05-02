package response

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
)

const (
	ErrBadRequestMessage     = "BAD_REQUEST"           // 400
	ErrNotFoundMessage       = "NOT_FOUND"             // 404
	ErrInternalServerMessage = "INTERNAL_SERVER_ERROR" // 500
	ErrUnauthorizedMessage   = "UNAUTHORIZED"
)

type HttpWriter interface {
	WriteCompress(data interface{})
	WriteError(err error, action string, msg ...string)
	WriteErrorBadRequest(err error, action string, params interface{})
	WriteErrorUnauthorized(err error, action string, params interface{})
}

type httpWriter struct {
	w http.ResponseWriter
	r *http.Request
}

func NewHttpWriter(w http.ResponseWriter, r *http.Request) HttpWriter {
	return &httpWriter{
		w: w,
		r: r,
	}
}

func (wt *httpWriter) WriteCompress(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		wt.WriteError(
			errors.WithStack(err),
			"ResponseJsonWriteCompress",
			"json parse error",
		)
		return
	}

	// When an error occurs in the compress process, should I change it to return uncompressed json?
	if len(jsonData) >= 2048 {
		if strings.Contains(wt.r.Header.Get("Accept-Encoding"), "gzip") {
			gz, err := gzip.NewWriterLevel(wt.w, gzip.DefaultCompression)
			if err != nil {
				wt.WriteError(
					errors.WithStack(err),
					"ResponseJsonWriteCompress",
					"gzip compress error",
				)
				return
			}
			defer gz.Close()
			wt.w.Header().Set("Content-Encoding", "gzip")
			wt.w.WriteHeader(http.StatusOK)
			gz.Write(jsonData)
			return
		}
		if strings.Contains(wt.r.Header.Get("Accept-Encoding"), "deflate") {
			df, err := flate.NewWriter(wt.w, flate.DefaultCompression)
			if err != nil {
				wt.WriteError(
					errors.WithStack(err),
					"ResponseJsonWriteCompress",
					"dfalte compress error",
				)
				return
			}
			defer df.Close()
			wt.w.Header().Set("Content-Encoding", "deflate")
			wt.w.WriteHeader(http.StatusOK)
			df.Write(jsonData)
			return
		}
	}

	wt.w.WriteHeader(http.StatusOK)
	wt.w.Write(jsonData)
}

func (wt *httpWriter) WriteError(err error, action string, msg ...string) {
	status := http.StatusInternalServerError
	message := ErrInternalServerMessage

	if err == sql.ErrNoRows {
		status = http.StatusNotFound
		message = ErrNotFoundMessage
	}

	wt.w.WriteHeader(status)
	json.NewEncoder(wt.w).Encode(
		map[string]interface{}{
			"status":  status,
			"message": message,
		},
	)

	var b bytes.Buffer
	if len(msg) > 0 {
		for _, v := range msg {
			b.WriteString(v)
		}
	}

	logger.NewDefaultLogger().
		Err(err).
		Str("Action", action).
		Msg(b.String())
}

func (wt *httpWriter) WriteErrorBadRequest(err error, action string, params interface{}) {
	status := http.StatusBadRequest

	wt.w.WriteHeader(status)
	json.NewEncoder(wt.w).Encode(
		map[string]interface{}{
			"status":  status,
			"message": ErrBadRequestMessage,
		},
	)

	logger.NewDefaultLogger().
		Err(err).
		Str("Action", action).
		Msgf("Params: %+v", params)
}

func (wt *httpWriter) WriteErrorUnauthorized(err error, action string, params interface{}) {
	status := http.StatusUnauthorized

	wt.w.WriteHeader(status)
	json.NewEncoder(wt.w).Encode(
		map[string]interface{}{
			"status":  status,
			"message": ErrUnauthorizedMessage,
		},
	)

	logger.NewDefaultLogger().
		Err(err).
		Str("Action", action).
		Msgf("Params: %+v", params)
}
