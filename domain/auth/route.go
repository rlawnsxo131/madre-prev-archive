package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"

	"github.com/rlawnsxo131/madre-server-v2/domain/auth/socialaccount"
	"github.com/rlawnsxo131/madre-server-v2/domain/user"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/lib/social"
	"github.com/rlawnsxo131/madre-server-v2/lib/token"

	"github.com/rlawnsxo131/madre-server-v2/utils"
)

func ApplyRoutes(v1 *mux.Router) {
	route := v1.NewRoute().PathPrefix("/auth").Subrouter()

	route.HandleFunc("", get()).Methods("GET")
	route.HandleFunc("", delete()).Methods("DELETE", "OPTIONS")
	route.HandleFunc("/google/check", postGoogleCheck()).Methods("POST", "OPTIONS")
	route.HandleFunc("/google/sign-in", postGoogleSignIn()).Methods("POST", "OPTIONS")
	route.HandleFunc("/google/sign-up", postGoogleSignUp()).Methods("POST", "OPTIONS")
}

func get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewHttpWriter(w, r)
		userTokenProfile := token.LoadUserTokenProfileFromHttpContextSyncMap(r.Context())

		writer.WriteCompress(
			map[string]interface{}{
				"user_token_profile": userTokenProfile,
			},
		)
	}
}

func delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewHttpWriter(w, r)

		token.ResetTokenCookies(w)

		writer.WriteCompress(
			map[string]bool{
				"is_success": true,
			},
		)
	}
}

func postGoogleCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewHttpWriter(w, r)
		db, err := database.LoadDBFromHttpSyncMapContext(r.Context())
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

		err = utils.NewValidator().Struct(&params)
		if err != nil {
			writer.WriteErrorBadRequest(
				err,
				"post /auth/google/check",
				&params,
			)
			return
		}

		profile, err := social.GetGoogleProfile(params.AccessToken)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/check",
			)
			return
		}

		// if no rows in result set err -> { exist: false }
		socialAccountService := socialaccount.NewService(db)
		socialAccount, err := socialAccountService.FindOneByProviderWithSocialId(
			socialaccount.Key_Provider_GOOGLE,
			profile.SocialId,
		)
		authService := NewService()
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

func postGoogleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewHttpWriter(w, r)
		db, err := database.LoadDBFromHttpSyncMapContext(r.Context())
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-in",
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
				"post /auth/google/sign-in",
				"decode params error",
			)
			return
		}

		err = utils.NewValidator().Struct(&params)
		if err != nil {
			writer.WriteErrorBadRequest(
				err,
				"post /auth/google/sign-in",
				&params,
			)
			return
		}

		profile, err := social.GetGoogleProfile(params.AccessToken)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		socialAccountService := socialaccount.NewService(db)
		socialAccount, err := socialAccountService.FindOneByProviderWithSocialId(
			socialaccount.Key_Provider_GOOGLE,
			profile.SocialId,
		)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		userService := user.NewService(db)
		user, err := userService.FindOneById(socialAccount.UserID)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		accessToken, refreshToken, err := token.GenerateTokens(
			user.ID,
			user.DisplayName,
			utils.NormalizeNullString(user.PhotoUrl),
		)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		token.SetTokenCookies(w, accessToken, refreshToken)

		writer.WriteCompress(map[string]interface{}{
			"user_token_profile": token.UserTokenProfile{
				DisplayName: user.DisplayName,
				PhotoUrl:    utils.NormalizeNullString(user.PhotoUrl),
				AccessToken: accessToken,
			},
		})
	}
}

func postGoogleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewHttpWriter(w, r)
		db, err := database.LoadDBFromHttpSyncMapContext(r.Context())
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-up",
			)
			return
		}

		var params struct {
			AccessToken string `json:"access_token" validate:"required,min=50"`
			DisplayName string `json:"display_name" validate:"required,max=16,min=1"`
		}

		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-up",
				"decode params error",
			)
			return
		}

		err = utils.NewValidator().Struct(&params)
		if err != nil {
			writer.WriteErrorBadRequest(
				err,
				"post /auth/google/sign-up",
				&params,
			)
			return
		}

		authService := NewService()
		valid, err := authService.ValidateDisplayName(params.DisplayName)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-up",
				"username validate error",
			)
			return
		}
		if !valid {
			writer.WriteErrorBadRequest(
				errors.New("username validation error"),
				"post /auth/google/sign-up",
				&params,
			)
			return
		}

		profile, err := social.GetGoogleProfile(params.AccessToken)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-up",
			)
			return
		}

		userService := user.NewService(db)
		userId, err := userService.Create(user.User{
			Email:       profile.Email,
			OriginName:  utils.NewNullString(profile.DisplayName),
			DisplayName: params.DisplayName,
			PhotoUrl:    utils.NewNullString(profile.PhotoUrl),
		})
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-up",
			)
			return
		}

		user, err := userService.FindOneById(userId)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-up",
			)
			return
		}

		socialAccountService := socialaccount.NewService(db)
		_, err = socialAccountService.Create(socialaccount.SocialAccount{
			UserID:   user.ID,
			Provider: "GOOGLE",
			SocialId: profile.SocialId,
		})
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-up",
			)
			return
		}

		accessToken, refreshToken, err := token.GenerateTokens(
			user.ID,
			user.DisplayName,
			utils.NormalizeNullString(user.PhotoUrl),
		)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		token.SetTokenCookies(w, accessToken, refreshToken)

		writer.WriteCompress(map[string]interface{}{
			"user_token_profile": token.UserTokenProfile{
				DisplayName: user.DisplayName,
				PhotoUrl:    utils.NormalizeNullString(user.PhotoUrl),
				AccessToken: accessToken,
			},
		})
	}
}
