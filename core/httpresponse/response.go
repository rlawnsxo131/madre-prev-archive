package httpresponse

import "github.com/rlawnsxo131/madre-server-v3/utils"

type response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message,omitempty"`
}

func NewResponse(code int, data any, message ...string) *response {
	return &response{
		Code:    code,
		Data:    data,
		Message: utils.ParseOptionalString(message...),
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

func NewErrorResponse(code int, strErr, message string) *errorResponse {
	return &errorResponse{
		Code:    code,
		Error:   strErr,
		Message: message,
	}
}
