package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v2/database"
)

func RegisterAPI(r chi.Router, db database.Database) {
	crudHandler := NewCRUDHandler(db)

	r.Route("/user", func(r chi.Router) {
		r.Get("/{id}", crudHandler.Get())
		r.Put("/{id}", crudHandler.Put())
	})
}
