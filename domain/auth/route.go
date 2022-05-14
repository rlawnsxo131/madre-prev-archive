package auth

import (
	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
)

func ApplyRoutes(r *mux.Router, db database.Database) {
	authRoute := r.NewRoute().PathPrefix("/auth").Subrouter()
	ctrl := NewController(db)

	authRoute.HandleFunc("", ctrl.Get()).Methods("GET")
	authRoute.HandleFunc("", ctrl.Delete()).Methods("DELETE", "OPTIONS")
	authRoute.HandleFunc("/google/check", ctrl.PostGoogleCheck()).Methods("POST", "OPTIONS")
	authRoute.HandleFunc("/google/sign-in", ctrl.PostGoogleSignIn()).Methods("POST", "OPTIONS")
	authRoute.HandleFunc("/google/sign-up", ctrl.PostGoogleSignUp()).Methods("POST", "OPTIONS")
}
