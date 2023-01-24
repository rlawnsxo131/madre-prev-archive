package apiv1

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v3/core/server/httpresponse"
	"github.com/rlawnsxo131/madre-server-v3/core/token"
)

type authRoute struct{}

func NewAuthRoute() *authRoute {
	return &authRoute{}
}

func (ar *authRoute) Register(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Delete("/", ar.Delete())
		r.Post("/google/check", ar.PostGoogleCheck())
	})
}

func (ar *authRoute) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)
		p := token.Profile(r.Context())

		if p == nil {
			rw.Error(
				errors.New("not found token profile"),
				httpresponse.NewErrorResponse(
					http.StatusUnauthorized,
				),
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
