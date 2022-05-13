package auth

import (
	"github.com/gorilla/mux"
)

func ApplyRoutes(v1 *mux.Router, ctrl Controller) {
	r := v1.NewRoute().PathPrefix("/auth").Subrouter()

	r.HandleFunc("", ctrl.Get()).Methods("GET")
	r.HandleFunc("", ctrl.Delete()).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/google/check", ctrl.PostGoogleCheck()).Methods("POST", "OPTIONS")
	r.HandleFunc("/google/sign-in", ctrl.PostGoogleSignIn()).Methods("POST", "OPTIONS")
	r.HandleFunc("/google/sign-up", ctrl.PostGoogleSignUp()).Methods("POST", "OPTIONS")
}
