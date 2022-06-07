package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v2/database"
)

func RegisterRoutes(r chi.Router, db database.Database) {
	ctrl := NewController(db)

	r.Route("/auth", func(r chi.Router) {
		r.Get("/", ctrl.Get())
		r.Delete("/", ctrl.Delete())
		r.Post("/google/check", ctrl.PostGoogleCheck())
		r.Post("/google/sign-in", ctrl.PostGoogleSignIn())
		r.Post("/google/sign-up", ctrl.PostGoogleSignUp())
	})
}
