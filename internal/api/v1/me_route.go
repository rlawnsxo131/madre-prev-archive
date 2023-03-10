package apiv1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/core/server/httpresponse"
	"github.com/rlawnsxo131/madre-server-v3/core/token"
	"github.com/rlawnsxo131/madre-server-v3/internal/application/handler/query"
)

type meRoute struct {
	userQueryhandler query.UserQueryHandler
}

func NewMeRoute() *meRoute {
	return &meRoute{}
}

func (mr *meRoute) Register(r chi.Router) {
	r.Route("/me", func(r chi.Router) {
		r.Get("/", mr.Get())
		r.Get("/info", mr.GetInfo())
	})
}

func (mr *meRoute) Get() http.HandlerFunc {
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

func (mr *meRoute) GetInfo() http.HandlerFunc {
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

		u, err := mr.userQueryhandler.Get(
			&query.GetUserQuery{
				UserId: p.UserId,
			},
		)

		if err != nil {
			rw.Error(
				err.Err,
				httpresponse.NewErrorResponse(
					err.Code,
					err.Message,
				),
			)
			return
		}

		rw.Write(
			httpresponse.NewResponse(
				http.StatusOK,
				u,
			),
		)
	}
}
