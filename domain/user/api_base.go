package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
)

type baseHandler struct {
	db database.Database
}

func NewBaseHandler(db database.Database) *baseHandler {
	return &baseHandler{
		db: db,
	}
}

func (h *baseHandler) Get() http.HandlerFunc {
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

func (h *baseHandler) Put() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
