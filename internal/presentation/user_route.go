package presentation

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/external/engine/httpresponse"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/command"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/query"
)

type userRoute struct {
	userService user.UserService
}

func NewUserRoute(db rdb.Database) *userRoute {
	return &userRoute{
		user.NewUserService(
			command.NewUserCommandRepository(db),
			query.NewUserQueryRepository(db),
		),
	}
}

func (h *userRoute) Register(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Get("/{id}", h.Get())
	})
}

func (h *userRoute) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)
		id := chi.URLParam(r, "id")

		u, err := h.userService.FindOneById(id)
		if err != nil {
			rw.Error(err)
			return
		}

		rw.Write(u)

	}
}
