package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v3/internal/datastore/rdb"
)

func RegisterHTTPHandler(r chi.Router, db rdb.Database) {
	handler := NewHTTPHandler(
		NewUserRepository(db),
	)

	r.Route("/user", func(r chi.Router) {
		r.Get("/{id}", handler.Get())
	})
}
