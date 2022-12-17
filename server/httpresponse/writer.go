package httpresponse

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/common"
	"github.com/rlawnsxo131/madre-server-v3/server/httplogger"
	"github.com/rs/zerolog"
)

type Writer interface {
	Write(res *response)
	Error(err error, msg ...string)
	ErrorBadRequest(err error, msg ...string)
	ErrorUnauthorized(err error, msg ...string)
	ErrorForbidden(err error, msg ...string)
	ErrorNotFound(err error, msg ...string)
	writeError(err error, code int, msg ...string)
}

type writer struct {
	w http.ResponseWriter
	r *http.Request
}

func NewWriter(w http.ResponseWriter, r *http.Request) Writer {
	return &writer{w, r}
}

func (wt *writer) Write(res *response) {
	jsonRes, err := json.Marshal(res)
	if err != nil {
		wt.Error(
			errors.Wrap(err, "Write json parse error"),
		)
		return
	}
	wt.w.WriteHeader(res.Code)
	wt.w.Write(jsonRes)
	httplogger.HTTPLoggerCtx(wt.r.Context()).Add(func(e *zerolog.Event) {
		e.RawJSON("response", jsonRes)
	})
}

func (wt *writer) Error(err error, msg ...string) {
	wt.writeError(
		err,
		parseError(err),
		msg...,
	)
}

func (wt *writer) ErrorBadRequest(err error, msg ...string) {
	wt.writeError(
		err,
		http.StatusBadRequest,
		msg...,
	)
}

func (wt *writer) ErrorUnauthorized(err error, msg ...string) {
	wt.writeError(
		err,
		http.StatusUnauthorized,
		msg...,
	)
}

func (wt *writer) ErrorForbidden(err error, msg ...string) {
	wt.writeError(
		err,
		http.StatusForbidden,
		msg...,
	)
}

func (wt *writer) ErrorNotFound(err error, msg ...string) {
	wt.writeError(
		err,
		http.StatusNotFound,
		msg...,
	)
}

func (wt *writer) writeError(err error, code int, msg ...string) {
	res, _ := json.Marshal(
		NewErrorResponse(
			code,
			msg...,
		),
	)
	wt.w.WriteHeader(code)
	wt.w.Write(res)
	httplogger.HTTPLoggerCtx(wt.r.Context()).Add(func(e *zerolog.Event) {
		e.Err(err).RawJSON("response", res)
	})
}

func parseError(err error) int {
	var code int

	switch {
	case errors.Is(err, sql.ErrNoRows):
		code = http.StatusNotFound

	case errors.Is(err, common.ErrMissingRequiredValue):
		code = http.StatusBadRequest

	case errors.Is(err, common.ErrNotSupportValue):
		code = http.StatusBadRequest

	case errors.Is(err, common.ErrConflictUniqValue):
		code = http.StatusConflict

	case errors.Is(err, common.ErrUnprocessableValue):
		code = http.StatusUnprocessableEntity

	default:
		code = http.StatusInternalServerError
	}

	return code
}
