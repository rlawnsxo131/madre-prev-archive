package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/lib"
)

func SetupRoute(v1 *mux.Router) {
	authRouter := v1.NewRoute().PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/google", getGoogle()).Methods("GET")
	authRouter.HandleFunc("/google/check/registerd", postGoogleRegisterd()).Methods("POST", "OPTIONS")
}

func getGoogle() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		auth := map[string]string{
			"google": "google",
		}
		lib.ResponseJsonCompressWriter(rw, r, auth)
	}
}

func postGoogleRegisterd() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var params postGoogleRegisterdParams
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			lib.ResponseErrorWriter(rw, err)
		}
		profile, err := lib.GetGoogleProfile(params.AccessToken)
		if err != nil {
			lib.ResponseErrorWriter(rw, err)
		}
		log.Println(profile)
		lib.ResponseJsonCompressWriter(rw, r, map[string]bool{"exist": false})
	}
}
