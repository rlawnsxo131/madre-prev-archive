package auth

import (
	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
)

func RegisterRoutes(r *mux.Router, db database.Database) {
	r = r.NewRoute().PathPrefix("/auth").Subrouter()
	ctrl := NewController(db)

	r.HandleFunc("", ctrl.Get()).Methods("GET")
	r.HandleFunc("", ctrl.Delete()).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/google/check", ctrl.PostGoogleCheck()).Methods("POST", "OPTIONS")
	r.HandleFunc("/google/sign-in", ctrl.PostGoogleSignIn()).Methods("POST", "OPTIONS")
	r.HandleFunc("/google/sign-up", ctrl.PostGoogleSignUp()).Methods("POST", "OPTIONS")
}
