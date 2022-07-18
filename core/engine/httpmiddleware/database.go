package httpmiddleware

import (
	"context"
	"net/http"

	"github.com/rlawnsxo131/madre-server-v3/core/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/core/engine/httpresponse"
)

func DatabasePool(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := rdb.Conn(r.Context())
		if err != nil {
			rw := httpresponse.NewWriter(w, r)
			rw.Error(err)
			return
		}
		defer conn.Release()

		dbCtx := context.WithValue(
			r.Context(),
			rdb.KEY_DATABASE_CTX,
			conn,
		)

		next.ServeHTTP(
			w,
			r.WithContext(dbCtx),
		)
	})
}
