package httpresponse

import "github.com/rlawnsxo131/madre-server-v3/utils"

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message,omitempty"`
}

func NewResponse(code int, data any, message ...string) *Response {
	return &Response{
		Code:    code,
		Data:    data,
		Message: utils.ParseOtionalString(message...),
	}
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

func NewErrorResponse(code int, strErr, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Error:   strErr,
		Message: message,
	}
}
