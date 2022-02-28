package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rlawnsxo131/madre-server-v2/constants"
	"github.com/urfave/negroni"
)

func CorsMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	allowHosts := []string{"http://localhost:8080", "http://localhost:5000"}
	originHeader := r.Header.Get("Origin")
	validation := false
	for _, host := range allowHosts {
		if originHeader == host {
			validation = true
		}
	}
	if validation {
		for _, method := range []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"} {
			if method == r.Method {
				rw.Header().Set("Access-Control-Allow-Origin", originHeader)
				rw.Header().Set("Access-Control-Allow-Credentials", "true")
				if method == "OPTIONS" {
					rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, Cookie, X-CSRF-Token")
					rw.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS,HEAD")
					return
				}
				break
			}
		}
	}
	next(rw, r)
}

func SetResponseContentTypeMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	next(rw, r)
}

func SetDatabaseContextMiddleware(db *sqlx.DB) negroni.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		ctx := context.WithValue(r.Context(), constants.DBContextKey, db)
		connectionCount := db.Stats().OpenConnections
		r.Context().Value(ctx)
		log.Println("connection count:", connectionCount)
		next(rw, r.WithContext(ctx))
	}
}
