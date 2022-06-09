package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v3/internal/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/internal/pkg/user"
)

func RegisterHTTPHandler(r chi.Router, db rdb.Database) {
	handler := NewHTTPHandler()
	googleHandler := NewHTTPGoogleHandler(
		NewSocialAccountRepository(db),
		user.NewUserRepository(db),
	)

	r.Route("/auth", func(r chi.Router) {
		r.Get("/", handler.Get())
		r.Delete("/", handler.Delete())
		r.Post("/google/check", googleHandler.PostGoogleCheck())
		r.Post("/google/sign-in", googleHandler.PostGoogleSignIn())
		r.Post("/google/sign-up", googleHandler.PostGoogleSignUp())
	})
}
