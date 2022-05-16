package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v2/database"
)

func RegisterRoutes(r chi.Router, db database.Database) {
	ctrl := NewController(db)
	r.Route("/user", func(r chi.Router) {
		r.Get("/{id}", ctrl.Get())
		r.Put("/{id}", ctrl.Put())
	})
}
