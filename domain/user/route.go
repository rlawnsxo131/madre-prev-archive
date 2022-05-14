package user

import (
	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
)

func ApplyRoutes(r *mux.Router, db database.Database) {
	userRoute := r.NewRoute().PathPrefix("/user").Subrouter()
	ctrl := NewController(db)

	userRoute.HandleFunc("/{id}", ctrl.Get()).Methods("GET")
	userRoute.HandleFunc("/{id}", ctrl.Put()).Methods("PUT", "OPTIONS")
}
