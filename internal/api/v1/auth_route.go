package apiv1

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v3/lib/token"
	"github.com/rlawnsxo131/madre-server-v3/server/httpresponse"
)

type authRoute struct{}

func NewAuthRoute() *authRoute {
	return &authRoute{}
}

func (ar *authRoute) Register(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Get("/", ar.Get())
		r.Delete("/", ar.Delete())
		r.Post("/google/check", ar.PostGoogleCheck())
	})
}

func (ar *authRoute) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)
		p := token.Profile(r.Context())

		rw.Write(
			httpresponse.NewResponse(
				http.StatusOK,
				p,
			),
		)
	}
}

func (ar *authRoute) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)
		p := token.Profile(r.Context())

		if p == nil {
			rw.ErrorUnauthorized(
				errors.New("not found token profile"),
			)
			return
		}
		token.NewManager().ResetCookies(w)

		rw.Write(
			httpresponse.NewResponse(
				http.StatusOK,
				nil,
			),
		)
	}
}

func (ar *authRoute) PostGoogleCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
