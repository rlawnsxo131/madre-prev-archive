package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib"
)

func SetupRoute(v1 *mux.Router) {
	authRouter := v1.NewRoute().PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/google", getGoogle()).Methods("GET")
	authRouter.HandleFunc("/google/check", postGoogleCheck()).Methods("POST")
	authRouter.HandleFunc("/google/signin", postGoogleSignin()).Methods("POST")
	authRouter.HandleFunc("/google/signup", postGoogleSignup()).Methods("POST")
}

func getGoogle() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		writer := lib.NewHttpWriter(rw, r)
		auth := map[string]string{
			"google": "google",
		}
		writer.WriteCompress(auth)
	}
}

func postGoogleCheck() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		writer := lib.NewHttpWriter(rw, r)
		db, err := database.GetDBConn(r.Context())
		if err != nil {
			writer.WriteError(err)
			return
		}

		var params struct {
			AccessToken string `json:"access_token"`
		}
		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			err = errors.Wrap(err, "post /auth/google/check: decode params error")
			writer.WriteError(err)
			return
		}

		googleProfileApi := lib.NewGooglePeopleApi(params.AccessToken)
		profile, err := googleProfileApi.GetGoogleProfile()
		if err != nil {
			writer.WriteError(err)
			return
		}

		// if no rows in result set err -> { exist: false }
		socialAccountService := NewSocialAccountService(db)
		socialAccount, err := socialAccountService.FindOneBySocialId(profile.SocialId)
		authService := NewAuthService()
		isExistSocialAccountMap, err := authService.GetIsExistSocialAccountMap(socialAccount, err)
		if err != nil {
			writer.WriteError(err)
			return
		}

		writer.WriteCompress(isExistSocialAccountMap)
	}
}

func postGoogleSignin() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {}
}

func postGoogleSignup() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {}
}
