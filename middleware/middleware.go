package middleware

import (
	"context"
	"errors"
	"runtime/debug"
	"time"

	"net/http"

	"sync"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rlawnsxo131/madre-server-v2/constants"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(
					w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError,
				)
				err := errors.New(string(debug.Stack()))
				logger.Logger.
					Err(err).
					Str("Action", "Recovery")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func HttpLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		httpLogger := logger.NewHttpLogger()

		bodyBuf, reader, err := httpLogger.ReadBody(r.Body)
		if err != nil {
			logger.Logger.
				Err(err).
				Str("Action", "HttpLogger ReadBody").
				Msgf("Params: %+v", r.Body)
		}
		r.Body = reader

		defer func() {
			httpLogger.LogEntry(r, start, string(bodyBuf))
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
