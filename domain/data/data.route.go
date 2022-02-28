package data

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib"
)

func SetupRoute(v1 *mux.Router) {
	dataRouter := v1.NewRoute().PathPrefix("/data").Subrouter()
	dataRouter.HandleFunc("", getAll()).Methods("GET")
	dataRouter.HandleFunc("/{uuid}", getOne()).Methods("GET")
}

func getAll() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			log.Printf("dataRoute: limit Atoi wrong: %v", err)
		}
		limit = lib.IfIsNotExistGetDefaultIntValue(limit, 50)

		db, err := database.GetDBConn(r.Context())
		if err != nil {
			lib.ResponseErrorWriter(rw, err)
			return
		}

		dataService := NewDataService(db)
		dataList, err := dataService.FindAll(limit)
		if err != nil {
			lib.ResponseErrorWriter(rw, err)
			return
		}

		lib.ResponseJsonCompressWriter(rw, r, dataList)
	}
}

func getOne() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		uuid := vars["uuid"]

		db, err := database.GetDBConn(r.Context())
		if err != nil {
			lib.ResponseErrorWriter(rw, err)
			return
		}

		dataService := NewDataService(db)
		data, err := dataService.FindOneByUUID(uuid)
		if err != nil {
			lib.ResponseErrorWriter(rw, err)
			return
		}

		lib.ResponseJsonCompressWriter(rw, r, data)
	}
}
