package data

import (
	"github.com/go-chi/chi/v5"

	"github.com/rlawnsxo131/madre-server-v2/database"
)

func RegisterRoutes(r chi.Router, db database.Database) {
	ctrl := NewController(db)
	r.Route("/data", func(r chi.Router) {
		r.Get("/", ctrl.GetAll())
		r.Get("/{id}", ctrl.Get())
	})
}
