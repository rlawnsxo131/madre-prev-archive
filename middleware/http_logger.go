package middleware

import (
	"net/http"
	"time"

	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
)

func HTTPLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hl := logger.NewHTTPLogger(r)

		t := time.Now()
		defer func() {
			hl.Write(t)
		}()

		next.ServeHTTP(w, RequestWithHTTPLogger(r, hl))
	})
}

func RequestWithHTTPLogger(r *http.Request, hl logger.HTTPLogger) *http.Request {
	r = r.WithContext(logger.SetHTTPLoggerCtx(r.Context(), hl))
	return r
}
