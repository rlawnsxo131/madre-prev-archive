package middleware

import (
	"net/http"
	"time"

	"github.com/rlawnsxo131/madre-server-v2/lib/httpcontext"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UTC()
		hl := logger.NewHTTPLogger(r)
		buf, err := hl.ReadBody(r)
		ctx := httpcontext.SetHTTPLogger(r.Context(), hl)
		if err != nil {
			rw := response.NewWriter(w, r)
			rw.Error(
				err,
				"HttpLogger",
			)
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
		hl.Write(start, string(buf))
	})
}
