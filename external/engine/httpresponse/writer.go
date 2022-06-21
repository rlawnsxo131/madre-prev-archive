package httpresponse

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/external/engine/httplogger"
	"github.com/rs/zerolog"
)

const (
	HTTP_CODE_BAD_REQUEST           = "BadRequest"          // 400
	HTTP_CODE_UNAUTHORIZED          = "Unauthorized"        // 401
	HTTP_CODE_FORBIDDEN             = "Forbidden"           // 403
	HTTP_CODE_NOT_FOUND             = "NotFound"            // 404
	HTTP_CODE_CONFLICT              = "Conflict"            // 409
	HTTP_CODE_UNPROCESSABLE_ENTITY  = "UnprocessableEntity" // 422
	HTTP_CODE_INTERNAL_SERVER_ERROR = "InternalServerError" // 500
)

type Writer interface {
	Write(data any)
	Error(err error)
	ErrorBadRequest(err error)
	ErrorUnauthorized(err error)
	ErrorForbidden(err error)
	ErrorNotFound(err error)
	ErrorConflict(err error)
	ErrorUnprocessableEntity(err error)
	standardError(status int, code string, err error)
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

func (wt *writer) Error(err error) {
	status := http.StatusInternalServerError
	code := HTTP_CODE_INTERNAL_SERVER_ERROR

	if err == sql.ErrNoRows {
		status = http.StatusNotFound
		code = HTTP_CODE_NOT_FOUND
	}

	wt.standardError(status, code, err)
}

func (wt *writer) ErrorBadRequest(err error) {
	wt.standardError(
		http.StatusBadRequest,
		HTTP_CODE_BAD_REQUEST,
		err,
	)
}

func (wt *writer) ErrorUnauthorized(err error) {
	wt.standardError(
		http.StatusUnauthorized,
		HTTP_CODE_UNAUTHORIZED,
		err,
	)
}

func (wt *writer) ErrorForbidden(err error) {
	wt.standardError(
		http.StatusForbidden,
		HTTP_CODE_FORBIDDEN,
		err,
	)
}

func (wt *writer) ErrorNotFound(err error) {
	wt.standardError(
		http.StatusNotFound,
		HTTP_CODE_NOT_FOUND,
		err,
	)
}

func (wt *writer) ErrorConflict(err error) {
	wt.standardError(
		http.StatusConflict,
		HTTP_CODE_CONFLICT,
		err,
	)
}

func (wt *writer) ErrorUnprocessableEntity(err error) {
	wt.standardError(
		http.StatusUnprocessableEntity,
		HTTP_CODE_UNPROCESSABLE_ENTITY,
		err,
	)
}

func (wt *writer) standardError(status int, code string, err error) {
	res, _ := json.Marshal(map[string]any{
		"status": status,
		"code":   code,
	})
	wt.w.WriteHeader(status)
	wt.w.Write(res)
	httplogger.LoggerCtx(wt.r.Context()).Add(func(e *zerolog.Event) {
		e.Err(err).RawJSON("response", res)
	})
}
