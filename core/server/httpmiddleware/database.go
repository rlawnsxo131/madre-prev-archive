package httpmiddleware

import (
	"net/http"

	"github.com/rlawnsxo131/madre-server-v3/core/datastore/rdb"
)

func Database(db rdb.SingletonDatabase) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dbCtx := rdb.SetDBCtx(
				r.Context(),
				db,
			)

			next.ServeHTTP(
				w,
				r.WithContext(dbCtx),
			)
		})
	}

}
