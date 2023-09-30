package httpresponse

import (
	"encoding/json"
	"net/http"

	"github.com/rlawnsxo131/madre-server/core/logger"
	"github.com/rs/zerolog"
)

type writer struct {
	w http.ResponseWriter
	r *http.Request
}

func NewWriter(w http.ResponseWriter, r *http.Request) *writer {
	return &writer{w, r}
}

func (wt *writer) JSON(res *response) {
	jsonRes, err := json.Marshal(res)
	if err != nil {
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

	wt.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	wt.w.WriteHeader(res.Code)
	wt.w.Write(jsonRes)

	logger.GetLogEntry(wt.r.Context()).Add(func(e *zerolog.Event) {
		e.RawJSON("response", jsonRes)
	})
}

func (wt *writer) ERROR(err error, res *errorResponse) {
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

	wt.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	wt.w.WriteHeader(code)
	wt.w.Write(jsonRes)

	logger.GetLogEntry(wt.r.Context()).Add(func(e *zerolog.Event) {
		e.Err(err).RawJSON("response", jsonRes)
		if parseErr != nil {
			e.Err(parseErr)
		}
	})
}
