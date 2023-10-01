package response

import (
	"net/http"
	"strings"
)

type httpResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data,omitempty"`
}

func NewHTTPResponse(code int, data any) *httpResponse {
	return &httpResponse{
		Code:   code,
		Status: http.StatusText(code),
		Data:   data,
	}
}

type httpErrorResponse struct {
	Code   int            `json:"code"`
	Status string         `json:"status"`
	Error  map[string]any `json:"error,omitempty"`
}

func NewHTTPErrorResponse(code int, data any, message ...string) *httpErrorResponse {
	return &httpErrorResponse{
		Code:   code,
		Status: http.StatusText(code),
		Error: map[string]any{
			"message": strings.Join(message, ""),
			"data":    data,
		},
	}
}
