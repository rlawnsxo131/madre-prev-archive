package httpresponse

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/core/engine/httplogger"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/common"
	"github.com/rlawnsxo131/madre-server-v3/utils"
	"github.com/rs/zerolog"
)

const (
	HTTP_ERROR_BAD_REQUEST           = "BadRequest"          // 400
	HTTP_ERROR_UNAUTHORIZED          = "Unauthorized"        // 401
	HTTP_ERROR_FORBIDDEN             = "Forbidden"           // 403
	HTTP_ERROR_NOT_FOUND             = "NotFound"            // 404
	HTTP_ERROR_CONFLICT              = "Conflict"            // 409
	HTTP_ERROR_UNPROCESSABLE_ENTITY  = "UnprocessableEntity" // 422
	HTTP_ERROR_INTERNAL_SERVER_ERROR = "InternalServerError" // 500
)

type Writer interface {
	Write(data any)
	Error(err error, message ...string)
	ErrorBadRequest(err error, message ...string)
	ErrorUnauthorized(err error, message ...string)
	ErrorForbidden(err error, message ...string)
	ErrorNotFound(err error, message ...string)
	writeError(code int, strErr string, err error, message ...string)
}

type writer struct {
	w http.ResponseWriter
	r *http.Request
}

func NewWriter(w http.ResponseWriter, r *http.Request) Writer {
	return &writer{w, r}
}

func (wt *writer) Write(data any) {
	res, err := json.Marshal(data)
	if err != nil {
		wt.Error(
			errors.Wrap(err, "Write json parse error"),
		)
		return
	}
	wt.w.WriteHeader(http.StatusOK)
	wt.w.Write(res)
	httplogger.LoggerCtx(wt.r.Context()).Add(func(e *zerolog.Event) {
		e.RawJSON("response", res)
	})
}

func (wt *writer) Error(err error, message ...string) {
	code, strErr := parseError(err)
	wt.writeError(
		code,
		strErr,
		err,
		message...,
	)
}

func (wt *writer) ErrorBadRequest(err error, message ...string) {
	wt.writeError(
		http.StatusBadRequest,
		HTTP_ERROR_BAD_REQUEST,
		err,
		message...,
	)
}

func (wt *writer) ErrorUnauthorized(err error, message ...string) {
	wt.writeError(
		http.StatusUnauthorized,
		HTTP_ERROR_UNAUTHORIZED,
		err,
		message...,
	)
}

func (wt *writer) ErrorForbidden(err error, message ...string) {
	wt.writeError(
		http.StatusForbidden,
		HTTP_ERROR_FORBIDDEN,
		err,
		message...,
	)
}

func (wt *writer) ErrorNotFound(err error, message ...string) {
	wt.writeError(
		http.StatusNotFound,
		HTTP_ERROR_NOT_FOUND,
		err,
		message...,
	)
}

func (wt *writer) writeError(
	code int,
	strErr string,
	err error,
	message ...string,
) {
	res, _ := json.Marshal(
		NewErrorResponse(
			code,
			strErr,
			utils.ParseOtionalString(message...),
		),
	)
	wt.w.WriteHeader(code)
	wt.w.Write(res)
	httplogger.LoggerCtx(wt.r.Context()).Add(func(e *zerolog.Event) {
		e.Err(err).RawJSON("response", res)
	})
}

func parseError(err error) (int, string) {
	var code int
	var strErr string

	switch {
	case errors.Is(err, sql.ErrNoRows):
		code = http.StatusNotFound
		strErr = HTTP_ERROR_NOT_FOUND

	case errors.Is(err, common.ErrMissingRequiredValue):
		code = http.StatusBadRequest
		strErr = HTTP_ERROR_BAD_REQUEST

	case errors.Is(err, common.ErrNotSupportValue):
		code = http.StatusBadRequest
		strErr = HTTP_ERROR_BAD_REQUEST

	case errors.Is(err, common.ErrConflictUniqValue):
		code = http.StatusConflict
		strErr = HTTP_ERROR_CONFLICT

	case errors.Is(err, common.ErrUnProcessableValue):
		code = http.StatusUnprocessableEntity
		strErr = HTTP_ERROR_UNPROCESSABLE_ENTITY

	default:
		code = http.StatusInternalServerError
		strErr = HTTP_ERROR_INTERNAL_SERVER_ERROR
	}

	return code, strErr
}
