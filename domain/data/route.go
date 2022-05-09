package data

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/lib/httpcontext"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

func ApplyRoutes(v1 *mux.Router) {
	route := v1.NewRoute().PathPrefix("/data").Subrouter()

	route.HandleFunc("", getAll()).Methods("GET")
	route.HandleFunc("/{id}", get()).Methods("GET")
}

func getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
		db, err := httpcontext.NewManager(r.Context()).Database()
		if err != nil {
			rw.Error(err, "get /data")
			return
		}

		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			logger.GetDefaultLogger().
				Warn().Msgf("route: limit Atoi wrong: %v", err)
		}
		limit = utils.IfIsNotExistGetDefaultIntValue(limit, 50)

		dataUseCase := NewUseCase(db)
		dd, err := dataUseCase.FindAll(limit)
		if err != nil {
			rw.Error(err, "get /data")
			return
		}

		rw.Compress(dd)
	}
}

func get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
		db, err := httpcontext.NewManager(r.Context()).Database()
		if err != nil {
			rw.Error(err, "get /data/{id}")
			return
		}

		vars := mux.Vars(r)
		id := vars["id"]

		dataUseCase := NewUseCase(db)
		d, err := dataUseCase.FindOneById(id)
		if err != nil {
			rw.Error(err, "get /data/{id}")
			return
		}

		rw.Compress(d)
	}
}
