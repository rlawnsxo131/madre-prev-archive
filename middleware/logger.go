package middleware

import (
	"net/http"
	"time"

	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		hl := logger.NewHttpLogger()
		buf, err := hl.ReadBody(r)
		if err != nil {
			rw := response.NewWriter(w, r)
			rw.Error(
				err,
				"HttpLogger",
			)
			return
		}
		next.ServeHTTP(w, r)
		hl.LogEntry(r, start, string(buf))
	})
}
