package data

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

func ApplyRoutes(v1 *mux.Router) {
	dataRoute := v1.NewRoute().PathPrefix("/data").Subrouter()

	dataRoute.HandleFunc("", getAll()).Methods("GET")
	dataRoute.HandleFunc("/{id}", get()).Methods("GET")
}

func getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewHttpWriter(w, r)
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			log.Printf("dataRoute: limit Atoi wrong: %v", err)
		}
		limit = utils.IfIsNotExistGetDefaultIntValue(limit, 50)

		db, err := database.GetDBConn(r.Context())
		if err != nil {
			writer.WriteError(err, "get /data")
			return
		}

		dataService := NewDataService(db)
		dataList, err := dataService.FindAll(limit)
		if err != nil {
			writer.WriteError(err, "get /data")
			return
		}

		writer.WriteCompress(dataList)
	}
}

func get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewHttpWriter(w, r)
		vars := mux.Vars(r)
		id := vars["id"]

		db, err := database.GetDBConn(r.Context())
		if err != nil {
			writer.WriteError(err, "get /data/{id}")
			return
		}

		dataService := NewDataService(db)
		data, err := dataService.FindOneById(id)
		if err != nil {
			writer.WriteError(err, "get /data/{id}")
			return
		}

		writer.WriteCompress(data)
	}
}
