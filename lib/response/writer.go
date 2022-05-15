package response

import (
	"compress/flate"
	"compress/gzip"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rs/zerolog"
)

const (
	Http_Msg_BadRequest          = "BadRequest"          // 400
	Http_Msg_Unauthorized        = "Unauthorized"        // 401
	Http_Msg_Forbidden           = "Forbidden"           // 403
	Http_Msg_NotFound            = "NotFound"            // 404
	Http_Msg_Conflict            = "Conflict"            // 409
	Http_Msg_InternalServerError = "InternalServerError" // 500
)

type Writer interface {
	Compress(data interface{})
	Error(err error)
	ErrorBadRequest(err error)
	ErrorUnauthorized(err error)
	ErrorForbidden(err error)
	ErrorConflict(err error)
	standardError(status int, msg string, err error)
}

type writer struct {
	w http.ResponseWriter
	r *http.Request
}

func NewWriter(w http.ResponseWriter, r *http.Request) Writer {
	return &writer{
		w: w,
		r: r,
	}
}

func (wt *writer) Compress(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		wt.Error(errors.Wrap(err, "compress json parse error"))
		return
	}

	// When an error occurs in the compress process, should I change it to return uncompressed json?
	if len(jsonData) >= 2048 {
		if strings.Contains(wt.r.Header.Get("Accept-Encoding"), "gzip") {
			gz, err := gzip.NewWriterLevel(wt.w, gzip.DefaultCompression)
			if err != nil {
				wt.Error(errors.Wrap(err, "gzip compress error"))
				return
			}
			defer gz.Close()
			wt.w.Header().Set("Content-Encoding", "gzip")
			wt.w.WriteHeader(http.StatusOK)
			gz.Write(jsonData)
			logger.HTTPLoggerCtx(wt.r.Context()).Add(func(e *zerolog.Event) {
				e.Int("status", http.StatusOK).RawJSON("response", jsonData)
			})
			return
		}
		if strings.Contains(wt.r.Header.Get("Accept-Encoding"), "deflate") {
			df, err := flate.NewWriter(wt.w, flate.DefaultCompression)
			if err != nil {
				wt.Error(errors.Wrap(err, "dfalte compress error"))
				return
			}
			defer df.Close()
			wt.w.Header().Set("Content-Encoding", "deflate")
			wt.w.WriteHeader(http.StatusOK)
			df.Write(jsonData)
			logger.HTTPLoggerCtx(wt.r.Context()).Add(func(e *zerolog.Event) {
				e.Int("status", http.StatusOK).RawJSON("response", jsonData)
			})
			return
		}
	}

	wt.w.WriteHeader(http.StatusOK)
	wt.w.Write(jsonData)
	logger.HTTPLoggerCtx(wt.r.Context()).Add(func(e *zerolog.Event) {
		e.Int("status", http.StatusOK).RawJSON("response", jsonData)
	})
}

func (wt *writer) Error(err error) {
	status := http.StatusInternalServerError
	message := Http_Msg_InternalServerError

	if err == sql.ErrNoRows {
		status = http.StatusNotFound
		message = Http_Msg_NotFound
	}

	wt.standardError(status, message, err)
}

func (wt *writer) ErrorBadRequest(err error) {
	wt.standardError(http.StatusBadRequest, Http_Msg_BadRequest, err)
}

func (wt *writer) ErrorUnauthorized(err error) {
	wt.standardError(http.StatusUnauthorized, Http_Msg_Unauthorized, err)
}

func (wt *writer) ErrorForbidden(err error) {
	wt.standardError(http.StatusForbidden, Http_Msg_Forbidden, err)
}

func (wt *writer) ErrorConflict(err error) {
	wt.standardError(http.StatusConflict, Http_Msg_Conflict, err)
}

func (wt *writer) standardError(status int, message string, err error) {
	jsonData, _ := json.Marshal(map[string]interface{}{
		"status":  status,
		"message": message,
	})

	wt.w.WriteHeader(status)
	wt.w.Write(jsonData)

	logger.HTTPLoggerCtx(wt.r.Context()).Add(func(e *zerolog.Event) {
		e.Err(err).Int("status", status).RawJSON("response", jsonData)
	})
}
