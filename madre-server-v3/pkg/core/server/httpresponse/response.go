package httpresponse

import (
	"net/http"

	valueutil "github.com/rlawnsxo131/madre-server-v3/pkg/core/utils/value-util"
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
	m := map[string]any{}
	if len(msg) > 0 {
		m["message"] = valueutil.ParseOptionalString(msg...)
	}

	return &errorResponse{
		Code:   code,
		Status: http.StatusText(code),
		Error:  m,
	}
}
