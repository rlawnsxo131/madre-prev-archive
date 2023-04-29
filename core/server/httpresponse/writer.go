package httpresponse

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/core/server/httplogger"
	"github.com/rs/zerolog"
)

type writer struct {
	w http.ResponseWriter
	r *http.Request
}

func NewWriter(w http.ResponseWriter, r *http.Request) *writer {
	return &writer{w, r}
}

func (wt *writer) Write(res *response) {
	jsonRes, err := json.Marshal(res)

	if err != nil {
		wt.Error(
			errors.Wrap(err, "Write json parse error"),
			NewErrorResponse(
				http.StatusInternalServerError,
			),
		)
		return
	}

	wt.w.WriteHeader(res.Code)
	wt.w.Write(jsonRes)

	if entry, err := httplogger.LogEntry(wt.r.Context()); err != nil {
		entry.Add(func(e *zerolog.Event) {
			e.RawJSON("response", jsonRes)
		})
		return
	}

	// TODO force logging
}

func (wt *writer) Error(err error, res *errorResponse) {
	jsonRes, parseErr := json.Marshal(res)

	if parseErr != nil {
		jsonRes = json.RawMessage(
			`{"code": 500, "status": "Internal Server Error", "error": {"message": "response json parse error"}}`,
		)
		wt.w.WriteHeader(http.StatusInternalServerError)
	} else {
		wt.w.WriteHeader(res.Code)
	}

	wt.w.Write(jsonRes)

	if entry, err := httplogger.LogEntry(wt.r.Context()); err != nil {
		entry.Add(func(e *zerolog.Event) {
			e.Err(err).RawJSON("response", jsonRes)
			if parseErr != nil {
				e.Err(parseErr)
			}
		})
		return
	}

	// TODO force logging
}
