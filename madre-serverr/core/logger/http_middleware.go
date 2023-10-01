package logger

import (
	"net/http"
	"time"

	chi_middleware "github.com/go-chi/chi/v5/middleware"

	"github.com/rlawnsxo131/madre-server/core/errorz"
)

func HTTPMiddleware(hl *HTTPLogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t := time.Now()
			ww := chi_middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			le := hl.NewLogEntry(r, ww)
			defer le.Write(t)

			err := le.ReadBody()
			if err != nil {
				ww.WriteHeader(http.StatusInternalServerError)
				ww.Write(
					[]byte(errorz.New(err).Error()),
				)
				return
			}

			next.ServeHTTP(
				ww,
				r.WithContext(
					SetLogEntry(r.Context(), le),
				),
			)
		})
	}
}
