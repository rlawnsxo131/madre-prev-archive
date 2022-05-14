package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
)

type Controller interface {
	Get() http.HandlerFunc
	Put() http.HandlerFunc
}

type controller struct {
	db database.Database
}

func NewController(db database.Database) Controller {
	return &controller{
		db: db,
	}
}

func (c *controller) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
		vars := mux.Vars(r)
		id := vars["id"]

		userUseCase := NewUseCase(c.db)
		user, err := userUseCase.FindOneById(id)
		if err != nil {
			rw.Error(err)
			return
		}

		rw.Compress(user)
	}
}

func (c *controller) Put() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
