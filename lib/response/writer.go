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
)

type HttpWriter interface {
	WriteCompress(data interface{})
	WriteError(err error, action string, msg ...string)
	WriteErrorBadRequest(err error, action string, params interface{})
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

func (writer *httpWriter) WriteCompress(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		writer.WriteError(
			errors.WithStack(err),
			"ResponseJsonWriteCompress",
			"json parse error",
		)
		return
	}

	// When an error occurs in the compress process, should I change it to return uncompressed json?
	if len(jsonData) >= 2048 {
		if strings.Contains(writer.r.Header.Get("Accept-Encoding"), "gzip") {
			gz, err := gzip.NewWriterLevel(writer.w, gzip.DefaultCompression)
			if err != nil {
				writer.WriteError(
					errors.WithStack(err),
					"ResponseJsonWriteCompress",
					"gzip compress error",
				)
				return
			}
			defer gz.Close()
			writer.w.Header().Set("Content-Encoding", "gzip")
			writer.w.WriteHeader(http.StatusOK)
			gz.Write(jsonData)
			return
		}
		if strings.Contains(writer.r.Header.Get("Accept-Encoding"), "deflate") {
			df, err := flate.NewWriter(writer.w, flate.DefaultCompression)
			if err != nil {
				writer.WriteError(
					errors.WithStack(err),
					"ResponseJsonWriteCompress",
					"dfalte compress error",
				)
				return
			}
			defer df.Close()
			writer.w.Header().Set("Content-Encoding", "deflate")
			writer.w.WriteHeader(http.StatusOK)
			df.Write(jsonData)
			return
		}
	}

	writer.w.WriteHeader(http.StatusOK)
	writer.w.Write(jsonData)
}

func (writer *httpWriter) WriteError(err error, action string, msg ...string) {
	status := http.StatusInternalServerError
	message := ErrInternalServerMessage

	if err == sql.ErrNoRows {
		status = http.StatusNotFound
		message = ErrNotFoundMessage
	}

	writer.w.WriteHeader(status)
	json.NewEncoder(writer.w).Encode(
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

	logger.Logger.
		Err(err).
		Str("Action", action).
		Msg(b.String())
}

func (writer *httpWriter) WriteErrorBadRequest(err error, action string, params interface{}) {
	status := http.StatusBadRequest

	writer.w.WriteHeader(status)
	json.NewEncoder(writer.w).Encode(
		map[string]interface{}{
			"status":  status,
			"message": ErrBadRequestMessage,
		},
	)

	logger.Logger.
		Err(err).
		Str("Action", action).
		Msgf("Params: %+v", params)
}
