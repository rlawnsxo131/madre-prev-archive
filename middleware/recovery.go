package middleware

import (
	"errors"
	"net/http"
	"runtime/debug"

	"github.com/rlawnsxo131/madre-server-v2/lib/response"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				writer := response.NewHttpWriter(w, r)
				writer.WriteError(
					errors.New(string(debug.Stack())),
					"Recovery",
				)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
