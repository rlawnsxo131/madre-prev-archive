package user

import (
	"github.com/gorilla/mux"
)

func ApplyRoutes(v1 *mux.Router, ctrl Controller) {
	r := v1.NewRoute().PathPrefix("/user").Subrouter()

	r.HandleFunc("/{id}", ctrl.Get()).Methods("GET")
	r.HandleFunc("/{id}", ctrl.Put()).Methods("PUT", "OPTIONS")
}
