package httpmiddleware

import (
	"net/http"

	"github.com/rlawnsxo131/madre-server-v3/core/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/core/server/httpresponse"
)

func DatabasePool(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := rdb.DBInstance()

		if err != nil {
			httpresponse.NewWriter(w, r).Error(
				err,
				httpresponse.NewErrorResponse(
					http.StatusInternalServerError,
				),
			)
			return
		}

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
