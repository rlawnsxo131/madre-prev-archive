package data

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type Controller interface {
	GetAll() http.HandlerFunc
	Get() http.HandlerFunc
}

type controller struct {
	db database.Database
}

func NewController(db database.Database) Controller {
	return &controller{
		db: db,
	}
}

func (c *controller) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			logger.GetDefaultLogger().
				Warn().Msgf("route: limit Atoi wrong: %v", err)
		}
		limit = utils.IfIsNotExistGetDefaultIntValue(limit, 50)

		dataUseCase := NewUseCase(c.db)
		dd, err := dataUseCase.FindAll(limit)
		if err != nil {
			rw.Error(err)
			return
		}

		rw.Compress(dd)
	}
}

func (c *controller) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)

		vars := mux.Vars(r)
		id := vars["id"]

		dataUseCase := NewUseCase(c.db)
		d, err := dataUseCase.FindOneById(id)
		if err != nil {
			rw.Error(err)
			return
		}

		rw.Compress(d)
	}
}
