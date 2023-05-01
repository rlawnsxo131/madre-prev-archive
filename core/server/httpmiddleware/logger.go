package httpmiddleware

import (
	"encoding/json"
	"net/http"
	"time"

	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rlawnsxo131/madre-server-v3/core/logger"
	"github.com/rlawnsxo131/madre-server-v3/core/server/httplogger"
	"github.com/rlawnsxo131/madre-server-v3/core/server/httpresponse"
	"github.com/rs/zerolog"
)

func Logger(hl httplogger.HTTPLogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t := time.Now()
			ww := chi_middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			le := hl.NewLogEntry(r, ww)
			defer le.Write(t)

			err := le.ReadBody()
			if err != nil {
				res, _ := json.Marshal(
					httpresponse.NewErrorResponse(
						http.StatusInternalServerError,
					),
				)
				ww.WriteHeader(http.StatusInternalServerError)
				ww.Write(res)

				logger.DefaultLogger.NewLogEntry().Add(func(e *zerolog.Event) {
					e.Err(err)
				}).SendError()
				return
			}

			loggerCtx := httplogger.SetLogEntry(
				r.Context(),
				le,
			)

			next.ServeHTTP(
				ww,
				r.WithContext(loggerCtx),
			)
		})
	}
}
