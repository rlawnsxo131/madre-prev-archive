package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/google"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"

	"github.com/rlawnsxo131/madre-server-v2/utils"
)

func ApplyRoutes(v1 *mux.Router) {
	authRoute := v1.NewRoute().PathPrefix("/auth").Subrouter()

	authRoute.HandleFunc("/google/check", postGoogleCheck()).Methods("POST", "OPTIONS")
	authRoute.HandleFunc("/google/signin", postGoogleSignin()).Methods("POST", "OPTIONS")
	authRoute.HandleFunc("/google/signup", postGoogleSignup()).Methods("POST", "OPTIONS")
}

func postGoogleCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewHttpWriter(w, r)
		db, err := database.GetDBConn(r.Context())
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/check",
			)
			return
		}

		var params struct {
			AccessToken string `json:"access_token" validate:"required,min=50"`
		}

		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			writer.WriteError(
				errors.WithStack(err),
				"post /auth/google/check",
				"decode params error",
			)
			return
		}

		err = utils.Validator.Struct(&params)
		if err != nil {
			writer.WriteErrorBadRequest(
				err,
				"post /auth/google/check",
				&params,
			)
			return
		}

		googleProfileApi := google.NewGooglePeopleApi(params.AccessToken)
		profile, err := googleProfileApi.GetGoogleProfile()
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/check",
			)
			return
		}

		// if no rows in result set err -> { exist: false }
		socialAccountService := NewSocialAccountService(db)
		socialAccount, err := socialAccountService.FindOneBySocialId(profile.SocialId)
		authService := NewAuthService()
		existSocialAccountMap, err := authService.GetExistSocialAccountMap(socialAccount, err)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/check",
			)
			return
		}
		writer.WriteCompress(existSocialAccountMap)
	}
}

func postGoogleSignin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func postGoogleSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewHttpWriter(w, r)
		db, err := database.GetDBConn(r.Context())
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/signup",
			)
			return
		}

		var params struct {
			AccessToken string `json:"access_token" validate:"requried,min=50"`
			Username    string `json:"username" validate:"required,min=1,max=16"`
		}

		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			writer.WriteError(err, "post /auth/google/signup", "decode params error")
			return
		}

		err = utils.Validator.Struct(&params)
		if err != nil {
			writer.WriteErrorBadRequest(
				err,
				"post /auth/google/signup",
				&params,
			)
			return
		}

		authService := NewAuthService()
		valid, err := authService.ValidateUserName(params.Username)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/signup",
				"username validate error",
			)
			return
		}
		if !valid {
			writer.WriteErrorBadRequest(
				errors.New("username validation error"),
				"post /auth/google/signup",
				&params,
			)
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
			writer.WriteError(
				err,
				"post /auth/google/signup",
			)
			return
		}

		socialAccount, err := socialAccountService.FindOneById(lastInsertId)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/signup",
			)
			return
		}

		writer.WriteCompress(socialAccount)
	}
}
