package httpmiddleware

import (
	"encoding/json"
	"net/http"
	"time"

	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rlawnsxo131/madre-server-v3/lib/logger"
	"github.com/rlawnsxo131/madre-server-v3/server/httplogger"
	"github.com/rlawnsxo131/madre-server-v3/server/httpresponse"
	"github.com/rs/zerolog"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		ww := chi_middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		hl := httplogger.NewHTTPLogger(r, ww)
		defer hl.Write(t)

		err := hl.ReadBody()
		if err != nil {
			res, _ := json.Marshal(httpresponse.NewErrorResponse(
				http.StatusInternalServerError,
			))
			ww.WriteHeader(http.StatusInternalServerError)
			ww.Write(res)

			logger.NewDefaultLogger().Add(func(e *zerolog.Event) {
				e.Err(err)
			}).SendError()
			return
		}

		loggerCtx := httplogger.SetHTTPLoggerCtx(
			r.Context(),
			hl,
		)

		next.ServeHTTP(
			ww,
			r.WithContext(loggerCtx),
		)
	})
}
