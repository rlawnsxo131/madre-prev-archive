package presentation

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v3/external/engine/httpresponse"
	"github.com/rlawnsxo131/madre-server-v3/lib/token"
)

type authRoute struct{}

func NewAuthRoute() *authRoute {
	return &authRoute{}
}

func (h *authRoute) Register(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Get("/", h.Get())
		r.Delete("/", h.Delete())
	})
}

func (h *authRoute) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)
		p := token.UserProfileCtx(r.Context())

		rw.Write(p)
	}
}

func (h *authRoute) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)
		p := token.UserProfileCtx(r.Context())
		if p == nil {
			rw.ErrorUnauthorized(
				errors.New("not found UserProfile"),
			)
			return
		}
		token.NewManager().ResetCookies(w)

		rw.Write(struct{}{})
	}
}
