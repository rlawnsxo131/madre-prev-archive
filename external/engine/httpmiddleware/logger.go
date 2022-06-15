package httpmiddleware

import (
	"encoding/json"
	"net/http"
	"time"

	chi_middleware "github.com/go-chi/chi/v5/middleware"

	"github.com/rlawnsxo131/madre-server-v3/external/engine/httplogger"
	"github.com/rlawnsxo131/madre-server-v3/external/engine/httpresponse"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		ww := chi_middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		hl := httplogger.NewLogger(r, ww)
		defer hl.Write(t)

		err := hl.ReadBody()
		if err != nil {
			res, _ := json.Marshal(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": httpresponse.HTTP_MSG_INTERNAL_SERVER_ERROR,
			})
			ww.WriteHeader(http.StatusInternalServerError)
			ww.Write(res)
			return
		}

		next.ServeHTTP(
			ww,
			r.WithContext(
				httplogger.SetLoggerCtx(
					r.Context(),
					hl,
				),
			),
		)
	})
}

// func RequestWithHTTPLogger(r *http.Request, hl logger.HTTPLogger) *http.Request {
// 	r = r.WithContext(r, hl))
// 	return r
// }
