package middleware

import (
	"errors"
	"net/http"
	"runtime/debug"

	"github.com/rlawnsxo131/madre-server-v2/lib/response"
)

var (
	allowHosts                = []string{"http://localhost:8080", "http://localhost:5000"}
	allowHostsWithoutProtocol = []string{"localhost:8080", "localhost:5000"}
)

func AllowHost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestHost := r.Host
		validation := false
		for _, host := range allowHostsWithoutProtocol {
			if requestHost == host {
				validation = true
				break
			}
		}
		if !validation {
			writer := response.NewHttpWriter(w, r)
			writer.WriteErrorForbidden(
				errors.New("forbidden host"),
				"AllowHost",
				requestHost,
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

func ContentTypeToJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
