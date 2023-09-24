package middleware

import (
	"fmt"
	"net/http"

	"github.com/rlawnsxo131/madre-server/core/errorz"
)

func AllowHost(allowHostsWithoutProtocol []string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
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
				w.Write([]byte("forbidden host"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func Cors(allowHosts []string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
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
}

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				newErr := errorz.New(
					fmt.Errorf("panic recover: %v", err),
				)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(newErr.Error()))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
