package user

import (
	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
)

func RegisterController(r *mux.Router, db database.Database) {
	r = r.NewRoute().PathPrefix("/user").Subrouter()
	ctrl := NewController(db)

	r.HandleFunc("/{id}", ctrl.Get()).Methods("GET")
	r.HandleFunc("/{id}", ctrl.Put()).Methods("PUT", "OPTIONS")
}
