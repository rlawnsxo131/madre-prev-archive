package data

import (
	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
)

func ApplyRoutes(v1 *mux.Router, db database.Database) {
	ctrl := NewController(db)
	r := v1.NewRoute().PathPrefix("/data").Subrouter()

	r.HandleFunc("", ctrl.GetAll()).Methods("GET")
	r.HandleFunc("/{id}", ctrl.Get()).Methods("GET")
}
