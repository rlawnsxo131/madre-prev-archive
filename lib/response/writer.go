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
	Http_Msg_BadRequest          = "BadRequest"          // 400
	Http_Msg_Unauthorized        = "Unauthorized"        // 401
	Http_Msg_Forbidden           = "Forbidden"           // 403
	Http_Msg_NotFound            = "NotFound"            // 404
	Http_Msg_Conflict            = "Conflict"            // 409
	Http_Msg_InternalServerError = "InternalServerError" // 500
)

type Writer interface {
	Compress(data interface{})
	Error(err error, action string, msg ...string)
	ErrorBadRequest(err error, action string, params interface{})
	ErrorUnauthorized(err error, action string, params interface{})
	ErrorForbidden(err error, action string, params interface{})
	ErrorConflict(err error, action string, params interface{})
	standardError(
		status int,
		Msg string,
		err error,
		action string,
		params interface{},
	)
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
		wt.Error(
			errors.WithStack(err),
			"ResponseJsonCompress",
			"json parse error",
		)
		return
	}

	// When an error occurs in the compress process, should I change it to return uncompressed json?
	if len(jsonData) >= 2048 {
		if strings.Contains(wt.r.Header.Get("Accept-Encoding"), "gzip") {
			gz, err := gzip.NewWriterLevel(wt.w, gzip.DefaultCompression)
			if err != nil {
				wt.Error(
					errors.WithStack(err),
					"ResponseJsonCompress",
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
				wt.Error(
					errors.WithStack(err),
					"ResponseJsonCompress",
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

func (wt *writer) Error(err error, action string, msg ...string) {
	status := http.StatusInternalServerError
	Msg := Http_Msg_InternalServerError

	if err == sql.ErrNoRows {
		status = http.StatusNotFound
		Msg = Http_Msg_NotFound
	}

	wt.w.WriteHeader(status)
	json.NewEncoder(wt.w).Encode(
		map[string]interface{}{
			"status": status,
			"Msg":    Msg,
		},
	)

	var b bytes.Buffer
	if len(msg) > 0 {
		for _, v := range msg {
			b.WriteString(v)
		}
	}

	logger.GetDefaultLogger().
		Err(err).
		Str("Action", action).
		Msg(b.String())
}

func (wt *writer) ErrorBadRequest(err error, action string, params interface{}) {
	wt.standardError(
		http.StatusBadRequest,
		Http_Msg_BadRequest,
		err,
		action,
		params,
	)
}

func (wt *writer) ErrorUnauthorized(err error, action string, params interface{}) {
	wt.standardError(
		http.StatusUnauthorized,
		Http_Msg_Unauthorized,
		err,
		action,
		params,
	)
}

func (wt *writer) ErrorForbidden(err error, action string, params interface{}) {
	wt.standardError(
		http.StatusForbidden,
		Http_Msg_Forbidden,
		err,
		action,
		params,
	)
}

func (wt *writer) ErrorConflict(err error, action string, params interface{}) {
	wt.standardError(
		http.StatusConflict,
		Http_Msg_Conflict,
		err,
		action,
		params,
	)

}

func (wt *writer) standardError(status int, Msg string, err error, action string, params interface{}) {
	wt.w.WriteHeader(status)
	json.NewEncoder(wt.w).Encode(
		map[string]interface{}{
			"status": status,
			"Msg":    Msg,
		},
	)

	logger.GetDefaultLogger().
		Err(err).
		Str("Action", action).
		Msgf("Params: %+v", params)
}
