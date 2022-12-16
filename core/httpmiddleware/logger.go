package httpmiddleware

import (
	"encoding/json"
	"net/http"
	"time"

	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rlawnsxo131/madre-server-v3/core/httpresponse"
	"github.com/rlawnsxo131/madre-server-v3/core/logger"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		ww := chi_middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		hl := logger.NewHTTPLogger(r, ww)
		defer hl.Write(t)

		err := hl.ReadBody()
		if err != nil {
			res, _ := json.Marshal(map[string]any{
				"code":  http.StatusInternalServerError,
				"error": httpresponse.HTTP_ERROR_INTERNAL_SERVER_ERROR,
			})
			ww.WriteHeader(http.StatusInternalServerError)
			ww.Write(res)
			return
		}

		loggerCtx := logger.SetHTTPLoggerCtx(
			r.Context(),
			hl,
		)

		next.ServeHTTP(
			ww,
			r.WithContext(loggerCtx),
		)
	})
}
