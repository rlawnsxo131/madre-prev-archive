package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
)

type crudHandler struct {
	db database.Database
}

func NewCRUDHandler(db database.Database) *crudHandler {
	return &crudHandler{
		db: db,
	}
}

func (h *crudHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
		id := chi.URLParam(r, "id")

		userUseCase := NewUserUseCase(
			NewUserRepository(h.db),
		)
		u, err := userUseCase.FindOneById(id)
		if err != nil {
			rw.Error(err)
			return
		}

		rw.Write(u)
	}
}

func (h *crudHandler) Put() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
