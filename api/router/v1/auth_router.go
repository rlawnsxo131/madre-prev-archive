package routerv1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type authRouter struct{}

func NewAuthRouter(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/check-google", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("check-google route"))
		})
	})
}
