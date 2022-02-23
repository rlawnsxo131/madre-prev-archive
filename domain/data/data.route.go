package data

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib"
)

func SetupRoute(v1 *mux.Router) {
	dataRouter := v1.NewRoute().PathPrefix("/data").Subrouter()
	dataRouter.HandleFunc("", getAll()).Methods("GET")
	dataRouter.HandleFunc("/{id}", getOne()).Methods("GET")
}

func getAll() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		limit = lib.IfIsNotExistGetDefaultIntValue(limit, 50)
		db := database.GetDBConn(r.Context())
		dataService := NewDataService(db)
		dataList, err := dataService.FindAll(limit)
		if err != nil {
			lib.ResponseErrorWriter(rw, err)
		}
		lib.ResponseJsonCompressWriter(rw, r, dataList)
	}
}

func getOne() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		db := database.GetDBConn(r.Context())
		dataService := NewDataService(db)
		data, err := dataService.FindOne(id)
		if err != nil {
			lib.ResponseErrorWriter(rw, err)
		}
		lib.ResponseJsonCompressWriter(rw, r, data)
	}
}
