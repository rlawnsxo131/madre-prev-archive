package useradapter

import (
	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v3/datastore/rdb"
)

func RegisterAPI(r chi.Router, db rdb.Database) {
	baseHandler := NewBaseHandler(db)

	r.Route("/user", func(r chi.Router) {
		r.Get("/{id}", baseHandler.Get())
	})
}
