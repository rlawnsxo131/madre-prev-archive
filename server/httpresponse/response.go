package httpresponse

import (
	"net/http"
	"reflect"

	"github.com/rlawnsxo131/madre-server-v3/lib/utils"
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
	Error  map[string]any `json:"error"`
}

func NewErrorResponse(code int, msg ...string) *errorResponse {
	m := map[string]any{
		"message": utils.ParseOptionalString(msg...),
	}
	for k, v := range m {
		if reflect.ValueOf(v).IsZero() {
			delete(m, k)
		}
	}
	return &errorResponse{
		Code:   code,
		Status: http.StatusText(code),
		Error:  m,
	}
}
