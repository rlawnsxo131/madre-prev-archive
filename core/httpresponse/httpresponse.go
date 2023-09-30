package httpresponse

import (
	"net/http"
	"strings"
)

type response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data,omitempty"`
}

func New(code int, data any) *response {
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

func NewError(code int, data any, message ...string) *errorResponse {
	return &errorResponse{
		Code:   code,
		Status: http.StatusText(code),
		Error: map[string]any{
			"message": strings.Join(message, ""),
			"data":    data,
		},
	}
}
