package response

import (
	"encoding/json"
	"net/http"

	"github.com/rlawnsxo131/madre-server/core/logger"
	"github.com/rs/zerolog"
)

type HTTPResponseWriter struct {
	w http.ResponseWriter
	r *http.Request
}

func NewHTTPResponseWriter() *HTTPResponseWriter {
	return &HTTPResponseWriter{}
}

func (hrw *HTTPResponseWriter) Init(w http.ResponseWriter, r *http.Request) *HTTPResponseWriter {
	hrw.w = w
	hrw.r = r

	return hrw
}

func (rw *HTTPResponseWriter) Json(res *httpResponse) {
	rw.respond(res.Code, res, nil)
}

func (rw *HTTPResponseWriter) Error(err error, res *httpErrorResponse) {
	rw.respond(res.Code, res, err)
}

func (rw *HTTPResponseWriter) respond(code int, res interface{}, err error) {
	jsonRes, parseErr := json.Marshal(res)
	if parseErr != nil {
		code = http.StatusInternalServerError
		jsonRes = json.RawMessage(`{
            "code": 500,
            "status": "Internal Server Error",
            "error": {
                "message": "response json parse error"
            }
        }`)
	}

	rw.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.w.WriteHeader(code)
	rw.w.Write(jsonRes)

	logger.GetLogEntry(rw.r.Context()).Add(func(e *zerolog.Event) {
		if err != nil {
			e.Err(err)
		}
		if parseErr != nil {
			e.Err(parseErr)
		}
		e.RawJSON("response", jsonRes)
	})
}
