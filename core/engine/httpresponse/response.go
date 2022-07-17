package httpresponse

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
