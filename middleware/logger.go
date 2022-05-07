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

		bodyBuf, reader, err := hl.ReadBody(r.Body)
		if err != nil {
			writer := response.NewHttpWriter(w, r)
			writer.Error(
				err,
				"HttpLogger",
			)
			return
		}
		r.Body = reader

		defer func() {
			hl.LogEntry(r, start, string(bodyBuf))
		}()
		next.ServeHTTP(w, r)
	})
}
