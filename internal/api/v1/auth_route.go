package apiv1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/core/server/httpresponse"
	"github.com/rlawnsxo131/madre-server-v3/core/token"
)

type authRoute struct{}

func NewAuthRoute() *authRoute {
	return &authRoute{}
}

func (ar *authRoute) Register(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/google/check", ar.postGoogleCheck())
		r.Delete("/", ar.delete())
	})
}

func (ar *authRoute) postGoogleCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (ar *authRoute) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)
		p := token.ProfileFromCtx(r.Context())

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
				http.StatusNoContent,
				nil,
			),
		)
	}
}
