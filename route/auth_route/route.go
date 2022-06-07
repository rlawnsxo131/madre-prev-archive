package authroute

import (
	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v2/database"
	socialRepo "github.com/rlawnsxo131/madre-server-v2/domain_v2/auth/repository"
	userRepo "github.com/rlawnsxo131/madre-server-v2/domain_v2/user/repository"
)

func RegisterRoutes(r chi.Router, db database.Database) {
	ctrl := NewController()
	googleCtrl := NewGoogleController(
		socialRepo.NewSocialAccountRepository(db),
		userRepo.NewUserRepository(db),
	)

	r.Route("/auth", func(r chi.Router) {
		r.Get("/", ctrl.Get())
		r.Delete("/", ctrl.Delete())
		r.Post("/google/check", googleCtrl.PostGoogleCheck())
		r.Post("/google/sign-in", googleCtrl.PostGoogleSignIn())
		r.Post("/google/sign-up", googleCtrl.PostGoogleSignUp())
	})
}
