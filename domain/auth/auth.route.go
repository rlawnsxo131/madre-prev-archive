package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/google"
	"github.com/rlawnsxo131/madre-server-v2/lib/router"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

func SetupRoute(v1 *mux.Router) {
	authRouter := v1.NewRoute().PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/google/check", postGoogleCheck()).Methods("POST")
	authRouter.HandleFunc("/google/signin", postGoogleSignin()).Methods("POST")
	authRouter.HandleFunc("/google/signup", postGoogleSignup()).Methods("POST")
}

func postGoogleCheck() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		writer := router.NewHttpWriter(rw, r)
		db, err := database.GetDBConn(r.Context())
		if err != nil {
			writer.WriteError(err)
			return
		}

		var params struct {
			AccessToken string `json:"access_token" validate:"required,min=50"`
		}

		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			err = errors.Wrap(err, "post /auth/google/check: decode params error")
			writer.WriteError(err)
			return
		}

		err = utils.ValidateManager.Struct(&params)
		if err != nil {
			writer.WriteErrorBadRequest("post /auth/google/check: params validation error", &params)
			return
		}

		googleProfileApi := google.NewGooglePeopleApi(params.AccessToken)
		profile, err := googleProfileApi.GetGoogleProfile()
		if err != nil {
			writer.WriteError(err)
			return
		}

		// if no rows in result set err -> { exist: false }
		socialAccountService := NewSocialAccountService(db)
		socialAccount, err := socialAccountService.FindOneBySocialId(profile.SocialId)
		authService := NewAuthService()
		existSocialAccountMap, err := authService.GetExistSocialAccountMap(socialAccount, err)
		if err != nil {
			writer.WriteError(err)
			return
		}
		writer.WriteCompress(existSocialAccountMap)
	}
}

func postGoogleSignin() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {}
}

func postGoogleSignup() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		writer := router.NewHttpWriter(rw, r)
		db, err := database.GetDBConn(r.Context())
		if err != nil {
			writer.WriteError(err)
			return
		}

		var params struct {
			AccessToken string `json:"access_token" validate:"requried,min=50"`
			Username    string `json:"username" validate:"required,min=1,max=16"`
		}

		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			err = errors.Wrap(err, "post /auth/google/signup: decode params error")
			writer.WriteError(err)
			return
		}

		err = utils.ValidateManager.Struct(&params)
		if err != nil {
			writer.WriteErrorBadRequest("post /auth/google/signup: params validation error", &params)
			return
		}

		authService := NewAuthService()
		valid, err := authService.ValidateUserName(params.Username)
		if err != nil {
			writer.WriteError(err)
			return
		}
		if !valid {
			writer.WriteErrorBadRequest("post /auth/google/signup: username validate error", &params)
			return
		}

		socialAccountService := NewSocialAccountService(db)
		lastInsertId, err := socialAccountService.Create(CreateSocialAccountParams{
			UUID:        utils.GenerateUUIDString(),
			AccessToken: params.AccessToken,
			UserName:    params.Username,
			Provider:    "GOOGLE",
		})
		if err != nil {
			writer.WriteError(err)
			return
		}

		socialAccount, err := socialAccountService.FindOneById(lastInsertId)
		if err != nil {
			writer.WriteError(err)
			return
		}

		writer.WriteCompress(socialAccount)
	}
}
