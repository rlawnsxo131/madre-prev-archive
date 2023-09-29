package httpresponse

import (
	"net/http"
)

type response struct {
	Code   int            `json:"code"`
	Status string         `json:"status"`
	Data   any            `json:"data,omitempty"`
	Error  *errorResponse `json:"error,omitempty"`
}

type errorResponse struct {
	Code    int
	Message string
}

func New(code int, data any, err *errorResponse) *response {
	return &response{
		Code:   code,
		Status: http.StatusText(code),
		Data:   data,
		Error:  err,
	}
}

// jsonRes = json.RawMessage(
// 	`{
// 		"code": 500,
// 		"status": "Internal Server Error",
// 		"error": {
// 			"message": "response json parse error"
// 		}
// 	}`,
// )

// w.Header().Set("Content-Type", "application/json; charset=utf-8")
// w.WriteHeader(code)
// w.Write(jsonRes)

// logger.GetLogEntry(r.Context()).Add(func(e *zerolog.Event) {
// 	e.RawJSON("response", jsonRes)
// })
