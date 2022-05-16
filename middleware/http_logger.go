package middleware

import (
	"net/http"
	"time"

	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
)

func HTTPLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := chi_middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		hl := logger.NewHTTPLogger(r, ww)

		t := time.Now()
		defer func() {
			hl.Write(t)
		}()

		next.ServeHTTP(ww, RequestWithHTTPLogger(r, hl))
	})
}

func RequestWithHTTPLogger(r *http.Request, hl logger.HTTPLogger) *http.Request {
	r = r.WithContext(logger.SetHTTPLoggerCtx(r.Context(), hl))
	return r
}
