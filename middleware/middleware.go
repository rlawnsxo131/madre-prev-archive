package middleware

import (
	"context"
	"errors"
	"log"
	"runtime/debug"
	"time"

	"net/http"

	"sync"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rlawnsxo131/madre-server-v2/constants"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/lib/token"
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

func HttpLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		httpLogger := logger.NewHttpLogger()

		bodyBytes, reader, err := httpLogger.ReadBody(r.Body)
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
			httpLogger.LogEntry(r, start, string(bodyBytes))
		}()
		next.ServeHTTP(w, r)
	})
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowHosts := []string{"http://localhost:8080", "http://localhost:5000"}
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

// TODO: need to separate the func that sets the context and the function that puts the context value.
func SetHttpContextValues(db *sqlx.DB) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			syncMap := sync.Map{}
			syncMap.Store(constants.HttpContextDBKey, db)
			ctx := context.WithValue(
				r.Context(),
				constants.HttpContextKey,
				syncMap,
			)
			r.Context().Value(ctx)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func SetResponseContentTypeJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := r.Cookie("Access_token")
		if err != nil {
			if err != http.ErrNoCookie {
				writer := response.NewHttpWriter(w, r)
				writer.WriteError(
					err,
					"JwtMiddleware",
					"get Access_token error",
				)
			}
			return
		}

		if accessToken != nil {
			tokenManager := token.NewTokenManager()
			claims, err := tokenManager.DecodeToken(accessToken.Value)
			if err != nil {
				refreshToken, err := r.Cookie("Refresh_token")
				if err != nil {
					if err != http.ErrNoCookie {
						writer := response.NewHttpWriter(w, r)
						writer.WriteError(
							err,
							"JwtMiddleware",
							"get Refresh_token error",
						)
						return
					}
				}

				if refreshToken != nil {
					claims, err := tokenManager.DecodeToken(refreshToken.Value)
					if err == nil {
						log.Println("refresh_token:", claims)
					} else {
						log.Println("err:", err)
					}
				}
			}
			log.Println("access_token:", claims)
		}

		next.ServeHTTP(w, r)
	})
}
