package httpmiddleware

import (
	"net/http"
	"runtime/debug"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/external/engine/httpresponse"
)

var (
	allowHosts                = []string{"http://localhost:8080", "http://localhost:5000"}
	allowHostsWithoutProtocol = []string{"localhost:8080", "localhost:5000"}
)

func AllowHost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqHost := r.Host
		validation := false
		for _, host := range allowHostsWithoutProtocol {
			if reqHost == host {
				validation = true
				break
			}
		}
		if !validation {
			rw := httpresponse.NewWriter(w, r)
			rw.ErrorForbidden(
				errors.New("AllowHost forbidden host"),
			)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		validation := false
		for _, host := range allowHosts {
			if origin == host {
				validation = true
				break
			}
		}
		if validation {
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
				rw := httpresponse.NewWriter(w, r)
				err := errors.New(string(debug.Stack()))
				rw.Error(
					errors.Wrap(err, "Recovery"),
				)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func ContentTypeToJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
