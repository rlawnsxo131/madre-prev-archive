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
	HTTP_MSG_BAD_REQUEST           = "BadRequest"          // 400
	HTTP_MSG_UNAUTHORIZED          = "Unauthorized"        // 401
	HTTP_MSG_FORBIDDEN             = "Forbidden"           // 403
	HTTP_MSG_NOT_FOUND             = "NotFound"            // 404
	HTTP_MSG_CONFLICT              = "Conflict"            // 409
	HTTP_MSG_UNPROCESSABLE_ENTITY  = "UnprocessableEntity" // 422
	HTTP_MSG_INTERNAL_SERVER_ERROR = "InternalServerError" // 500
)

type Writer interface {
	Write(data interface{})
	Error(err error)
	ErrorBadRequest(err error)
	ErrorUnauthorized(err error)
	ErrorForbidden(err error)
	ErrorNotFound(err error)
	ErrorConflict(err error)
	ErrorUnprocessableEntity(err error)
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

func (wt *writer) Write(data interface{}) {
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
	message := HTTP_MSG_INTERNAL_SERVER_ERROR

	if err == sql.ErrNoRows {
		status = http.StatusNotFound
		message = HTTP_MSG_NOT_FOUND
	}

	wt.standardError(status, message, err)
}

func (wt *writer) ErrorBadRequest(err error) {
	wt.standardError(http.StatusBadRequest, HTTP_MSG_BAD_REQUEST, err)
}

func (wt *writer) ErrorUnauthorized(err error) {
	wt.standardError(http.StatusUnauthorized, HTTP_MSG_UNAUTHORIZED, err)
}

func (wt *writer) ErrorForbidden(err error) {
	wt.standardError(http.StatusForbidden, HTTP_MSG_FORBIDDEN, err)
}

func (wt *writer) ErrorNotFound(err error) {
	wt.standardError(http.StatusNotFound, HTTP_MSG_NOT_FOUND, err)
}

func (wt *writer) ErrorConflict(err error) {
	wt.standardError(http.StatusConflict, HTTP_MSG_CONFLICT, err)
}

func (wt *writer) ErrorUnprocessableEntity(err error) {
	wt.standardError(http.StatusUnprocessableEntity, HTTP_MSG_UNPROCESSABLE_ENTITY, err)
}

func (wt *writer) standardError(status int, message string, err error) {
	res, _ := json.Marshal(map[string]interface{}{
		"status":  status,
		"message": message,
	})
	wt.w.WriteHeader(status)
	wt.w.Write(res)
	httplogger.LoggerCtx(wt.r.Context()).Add(func(e *zerolog.Event) {
		e.Err(err).RawJSON("response", res)
	})
}
