package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v2/database"
)

func RegisterAPI(r chi.Router, db database.Database) {
	baseHandler := NewBaseHandler(db)

	r.Route("/user", func(r chi.Router) {
		r.Get("/{id}", baseHandler.Get())
		r.Put("/{id}", baseHandler.Put())
	})
}
