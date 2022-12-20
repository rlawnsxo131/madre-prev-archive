package httpmiddleware

import (
	"net/http"

	"github.com/rlawnsxo131/madre-server-v3/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/server/httpresponse"
)

func DatabasePool(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := rdb.Conn()
		if err != nil {
			rw := httpresponse.NewWriter(w, r)
			rw.Error(
				err,
				httpresponse.NewErrorResponse(
					http.StatusInternalServerError,
				),
			)
			return
		}
		defer conn.Release()

		dbCtx := rdb.SetConnCtx(
			r.Context(),
			conn,
		)

		next.ServeHTTP(
			w,
			r.WithContext(dbCtx),
		)
	})
}
