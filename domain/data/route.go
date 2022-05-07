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

func ApplyRoutes(v1 *mux.Router) {
	route := v1.NewRoute().PathPrefix("/data").Subrouter()

	route.HandleFunc("", getAll()).Methods("GET")
	route.HandleFunc("/{id}", get()).Methods("GET")
}

func getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewHttpWriter(w, r)
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			logger.NewDefaultLogger().
				Warn().Msgf("route: limit Atoi wrong: %v", err)
		}
		limit = utils.IfIsNotExistGetDefaultIntValue(limit, 50)

		db, err := database.LoadFromHttpCtx(r.Context())
		if err != nil {
			writer.Error(err, "get /data")
			return
		}

		dService := NewService(db)
		dd, err := dService.FindAll(limit)
		if err != nil {
			writer.Error(err, "get /data")
			return
		}

		writer.Compress(dd)
	}
}

func get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewHttpWriter(w, r)
		vars := mux.Vars(r)
		id := vars["id"]

		db, err := database.LoadFromHttpCtx(r.Context())
		if err != nil {
			writer.Error(err, "get /data/{id}")
			return
		}

		dService := NewService(db)
		d, err := dService.FindOneById(id)
		if err != nil {
			writer.Error(err, "get /data/{id}")
			return
		}

		writer.Compress(d)
	}
}
