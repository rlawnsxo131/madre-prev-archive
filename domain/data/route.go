package data

import (
	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
)

func ApplyRoutes(r *mux.Router, db database.Database) {
	dataRoute := r.NewRoute().PathPrefix("/data").Subrouter()
	ctrl := NewController(db)

	dataRoute.HandleFunc("", ctrl.GetAll()).Methods("GET")
	dataRoute.HandleFunc("/{id}", ctrl.Get()).Methods("GET")
}
