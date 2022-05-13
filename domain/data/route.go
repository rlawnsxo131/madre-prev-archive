package data

import (
	"github.com/gorilla/mux"
)

func ApplyRoutes(v1 *mux.Router, ctrl Controller) {
	r := v1.NewRoute().PathPrefix("/data").Subrouter()

	r.HandleFunc("", ctrl.GetAll()).Methods("GET")
	r.HandleFunc("/{id}", ctrl.Get()).Methods("GET")
}
