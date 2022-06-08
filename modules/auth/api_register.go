package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v2/database"
)

func RegisterAPI(r chi.Router, db database.Database) {
	baseHandler := NewBaseHandler(db)
	googleHandler := NewGoogleHandler(db)

	r.Route("/auth", func(r chi.Router) {
		r.Get("/", baseHandler.Get())
		r.Delete("/", baseHandler.Delete())
		r.Post("/google/check", googleHandler.PostGoogleCheck())
		r.Post("/google/sign-in", googleHandler.PostGoogleSignIn())
		r.Post("/google/sign-up", googleHandler.PostGoogleSignUp())
	})
}
