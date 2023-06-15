package apiv1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/core/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/core/server/httpresponse"
	"github.com/rlawnsxo131/madre-server-v3/core/token"
	"github.com/rlawnsxo131/madre-server-v3/internal/application/handler/query"
	queryrepository "github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/persistence/repository/query"
)

type meRoute struct {
	userQueryHandler query.UserQueryHandler
}

func NewMeRoute(db rdb.SingletonDatabase) *meRoute {
	return &meRoute{
		userQueryHandler: query.NewUserQueryHandler(
			queryrepository.NewUserQueryRepository(db),
			queryrepository.NewUserSocialAccountQueryRepository(db),
		),
	}
}

func (mr *meRoute) Register(r chi.Router) {
	r.Route("/me", func(r chi.Router) {
		r.Get("/", mr.get())
		r.Get("/info", mr.getInfo())
	})
}

func (mr *meRoute) get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := token.ProfileFromCtx(r.Context())

		httpresponse.NewWriter(w, r).Json(
			httpresponse.NewResponse(
				http.StatusOK,
				p,
			),
		)
	}
}

func (mr *meRoute) getInfo() http.HandlerFunc {
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

		u, err := mr.userQueryHandler.Get(
			&query.GetUserQuery{
				UserId: p.UserId,
			},
		)

		if err != nil {
			rw.Error(
				err.Err,
				httpresponse.NewError(
					err.Code,
					err.Message,
				),
			)
			return
		}

		rw.Json(
			httpresponse.NewResponse(
				http.StatusOK,
				u,
			),
		)
	}
}
