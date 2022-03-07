package data

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib"
)

var utils = lib.GetUtils()

func SetupRoute(v1 *mux.Router) {
	dataRouter := v1.NewRoute().PathPrefix("/data").Subrouter()
	dataRouter.HandleFunc("", getAll()).Methods("GET")
	dataRouter.HandleFunc("/{uuid}", get()).Methods("GET")
}

func getAll() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		writer := lib.NewHttpWriter(rw, r)
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			log.Printf("dataRoute: limit Atoi wrong: %v", err)
		}
		limit = utils.IfIsNotExistGetDefaultIntValue(limit, 50)

		db, err := database.GetDBConn(r.Context())
		if err != nil {
			writer.WriteError(err)
			return
		}

		dataService := NewDataService(db)
		dataList, err := dataService.FindAll(limit)
		if err != nil {
			writer.WriteError(err)
			return
		}

		writer.WriteCompress(dataList)
	}
}

func get() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		writer := lib.NewHttpWriter(rw, r)
		vars := mux.Vars(r)
		uuid := vars["uuid"]

		db, err := database.GetDBConn(r.Context())
		if err != nil {
			writer.WriteError(err)
			return
		}

		dataService := NewDataService(db)
		data, err := dataService.FindOneByUUID(uuid)
		if err != nil {
			writer.WriteError(err)
			return
		}

		writer.WriteCompress(data)
	}
}
