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

	entry, entryErr := httplogger.LogEntry(wt.r.Context())
	if entryErr != nil {
		// TODO force logging
		return
	}

	entry.Add(func(e *zerolog.Event) {
		e.RawJSON("response", jsonRes)
	})

}

func (wt *writer) Error(err error, res *errorResponse) {
	code := res.Code
	jsonRes, parseErr := json.Marshal(res)
	if parseErr != nil {
		code = http.StatusInternalServerError
		jsonRes = json.RawMessage(
			`{
				"code": 500, 
				"status": "Internal Server Error", 
				"error": {
					"message": "response json parse error"
				}
			}`,
		)
	}

	wt.w.WriteHeader(code)
	wt.w.Write(jsonRes)

	entry, entryErr := httplogger.LogEntry(wt.r.Context())
	if entryErr != nil {
		// TODO force logging
		return
	}

	entry.Add(func(e *zerolog.Event) {
		e.Err(err).RawJSON("response", jsonRes)
		if parseErr != nil {
			e.Err(parseErr)
		}
	})

}
