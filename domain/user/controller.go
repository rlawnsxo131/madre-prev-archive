package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
)

type controller struct {
	db database.Database
}

func NewController(db database.Database) *controller {
	return &controller{
		db: db,
	}
}

func (c *controller) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
		id := chi.URLParam(r, "id")

		userUseCase := NewUserUseCase(
			NewUserRepository(c.db),
		)
		u, err := userUseCase.FindOneById(id)
		if err != nil {
			rw.Error(err)
			return
		}

		rw.Write(u)
	}
}

func (c *controller) Put() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
