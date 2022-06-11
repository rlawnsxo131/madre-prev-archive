package presentation

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/external/engine/httpresponse"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/repository"
)

type userHandler struct {
	userService user.UserService
}

func NewUserHandler(db rdb.Database) *userHandler {
	return &userHandler{
		user.NewUserService(
			repository.NewUserCommandRepository(db),
			repository.NewUserQueryRepository(db),
		),
	}
}

func (h *userHandler) Register(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Get("/{id}", h.Get())
	})
}

func (h *userHandler) Get() http.HandlerFunc {
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
