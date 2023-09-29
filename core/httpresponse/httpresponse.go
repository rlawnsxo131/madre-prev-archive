package httpresponse

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/rlawnsxo131/madre-server/core/logger"
	"github.com/rs/zerolog"
)

type response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

func NewResponse(code int, data any) *response {
	return &response{
		Code:   code,
		Status: http.StatusText(code),
		Data:   data,
	}
}

type errorResponse struct {
	Code   int            `json:"code"`
	Status string         `json:"status"`
	Error  map[string]any `json:"error,omitempty"`
}

func NewError(code int, msg ...string) *errorResponse {
	return &errorResponse{
		Code:   code,
		Status: http.StatusText(code),
		Error: map[string]any{
			"message": strings.Join(msg, ""),
		},
	}
}

func Json(
	w http.ResponseWriter,
	r *http.Request,
	res *response,
) {
	code := res.Code
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

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(jsonRes)

	logger.GetLogEntry(r.Context()).Add(func(e *zerolog.Event) {
		e.RawJSON("response", jsonRes)
	})
}

func Error(
	w http.ResponseWriter,
	r *http.Request,
	err error,
	res *errorResponse,
) {

	code := res.Code
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

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(jsonRes)

	logger.GetLogEntry(r.Context()).Add(func(e *zerolog.Event) {
		e.Err(err).RawJSON("response", jsonRes)
	})
}
