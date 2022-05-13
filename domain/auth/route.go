package auth

import (
	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
)

func ApplyRoutes(v1 *mux.Router, db database.Database) {
	ctrl := NewController(db)
	r := v1.NewRoute().PathPrefix("/auth").Subrouter()

	r.HandleFunc("", ctrl.Get()).Methods("GET")
	r.HandleFunc("", ctrl.Delete()).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/google/check", ctrl.PostGoogleCheck()).Methods("POST", "OPTIONS")
	r.HandleFunc("/google/sign-in", ctrl.PostGoogleSignIn()).Methods("POST", "OPTIONS")
	r.HandleFunc("/google/sign-up", ctrl.PostGoogleSignUp()).Methods("POST", "OPTIONS")
}
