package middleware

import (
	"net/http"

	"github.com/rlawnsxo131/madre-server-v2/constants"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/lib/syncmap"
)

func SetSyncMapContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := syncmap.GenerateHttpContext(r.Context())
		r.Context().Value(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SetDatabaseContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := database.GetDatabase()
		if err != nil {
			writer := response.NewHttpWriter(w, r)
			writer.WriteError(
				err,
				"SetDBContext",
			)
			return
		}

		ctx, err := syncmap.SetNewValueFromHttpContext(
			r.Context(),
			constants.Key_HttpContextDB,
			db,
		)
		if err != nil {
			writer := response.NewHttpWriter(w, r)
			writer.WriteError(
				err,
				"SetDBContext",
				"context set error",
			)
			return
		}

		r.Context().Value(ctx)
		next.ServeHTTP(w, r)
	})
}
