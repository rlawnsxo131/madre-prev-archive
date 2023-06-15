package apiv1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/core/server/httpresponse"
	"github.com/rlawnsxo131/madre-server-v3/core/token"
	"github.com/rlawnsxo131/madre-server-v3/internal/application/handler/command"
)

type authRoute struct {
	userCommandHandler command.UserCommandHandler
}

func NewAuthRoute(r chi.Router, userCommandHandler command.UserCommandHandler) *authRoute {
	ar := &authRoute{
		userCommandHandler,
	}

	r.Route("/auth", func(r chi.Router) {
		r.Post("/google/check", ar.postGoogleCheck())
		r.Delete("/", ar.delete())
	})

	return ar
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
				httpresponse.NewError(
					http.StatusUnauthorized,
				),
			)
			return
		}

		token.NewManager().ResetCookies(w)

		rw.Json(
			httpresponse.NewResponse(
				http.StatusNoContent,
				nil,
			),
		)
	}
}
