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
		httpLogger := logger.NewHttpLogger()

		bodyBuffer, reader, err := httpLogger.ReadBody(r.Body)
		if err != nil {
			writer := response.NewHttpWriter(w, r)
			writer.WriteError(
				err,
				"HttpLogger",
			)
			return
		}
		r.Body = reader

		defer func() {
			httpLogger.LogEntry(r, start, string(bodyBuffer))
		}()
		next.ServeHTTP(w, r)
	})
}
