package httpmiddleware

import (
	"net/http"
	"runtime/debug"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/pkg/core/server/httpresponse"
)

var (
	allowHosts                = []string{"http://localhost:8080", "http://localhost:5000"}
	allowHostsWithoutProtocol = []string{"localhost:8080", "localhost:5000"}
)

func AllowHost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqHost := r.Host
		valid := false
		for _, host := range allowHostsWithoutProtocol {
			if reqHost == host {
				valid = true
				break
			}
		}
		if !valid {
			httpresponse.NewWriter(w, r).Error(
				errors.New("AllowHost forbidden host"),
				httpresponse.NewError(
					http.StatusForbidden,
				),
			)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		valid := false
		for _, host := range allowHosts {
			if origin == host {
				valid = true
				break
			}
		}
		if valid {
			for _, method := range []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"} {
				if method == r.Method {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					w.Header().Set("Access-Control-Allow-Credentials", "true")
					if method == "OPTIONS" {
						w.Header().Set(
							"Access-Control-Allow-Headers",
							"Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, Cookie, X-CSRF-Token",
						)
						w.Header().Set(
							"Access-Control-Allow-Methods",
							"GET,POST,PUT,PATCH,DELETE,OPTIONS,HEAD",
						)
						return
					}
					break
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				err := errors.Wrap(
					errors.New(string(debug.Stack())),
					"Recovery",
				)
				httpresponse.NewWriter(w, r).Error(
					err,
					httpresponse.NewError(
						http.StatusInternalServerError,
					),
				)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
