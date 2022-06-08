package useradapter

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v3/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/domain/user/userrepository"

	"github.com/rlawnsxo131/madre-server-v3/lib/response"
)

type baseHandler struct {
	db rdb.Database
}

func NewBaseHandler(db rdb.Database) *baseHandler {
	return &baseHandler{
		db: db,
	}
}

func (h *baseHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
		id := chi.URLParam(r, "id")

		userUseCase := user.NewUserUseCase(
			userrepository.NewUserRepository(h.db),
		)
		u, err := userUseCase.FindOneById(id)
		if err != nil {
			rw.Error(err)
			return
		}

		rw.Write(u)
	}
}
