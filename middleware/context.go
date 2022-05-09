package middleware

import (
	"net/http"

	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/httpcontext"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
)

func SetDatabaseCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := database.GetDatabaseInstance()
		if err != nil {
			rw := response.NewWriter(w, r)
			rw.Error(
				err,
				"SetDatabaseCtx",
			)
			return
		}
		cm := httpcontext.NewManager(r.Context())
		ctx := cm.SetDatabase(db)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
