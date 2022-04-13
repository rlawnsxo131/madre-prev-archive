package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/domain/user"
	"github.com/rlawnsxo131/madre-server-v2/lib/google"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/lib/syncmap"
	"github.com/rlawnsxo131/madre-server-v2/lib/token"

	"github.com/rlawnsxo131/madre-server-v2/utils"
)

func ApplyRoutes(v1 *mux.Router) {
	authRoute := v1.NewRoute().PathPrefix("/auth").Subrouter()

	authRoute.HandleFunc("", get()).Methods("GET")
	authRoute.HandleFunc("/sign-out", postSignOut()).Methods("POST", "OPTIONS")
	authRoute.HandleFunc("/google/check", postGoogleCheck()).Methods("POST", "OPTIONS")
	authRoute.HandleFunc("/google/sign-in", postGoogleSignIn()).Methods("POST", "OPTIONS")
	authRoute.HandleFunc("/google/sign-up", postGoogleSignUp()).Methods("POST", "OPTIONS")
}

func get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewHttpWriter(w, r)
		profile := syncmap.LoadUserTokenProfileFromHttpContext(r.Context())

		writer.WriteCompress(
			map[string]interface{}{
				"user_token_profile": profile,
			},
		)
	}
}

func postSignOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
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

		profile, err := google.GetPeopleProfile(params.AccessToken)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/check",
			)
			return
		}

		// if no rows in result set err -> { exist: false }
		socialAccountService := NewSocialAccountService(db)
		socialAccount, err := socialAccountService.FindOneByProviderWithSocialId(
			Key_Provider_GOOGLE,
			profile.SocialId,
		)
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

func postGoogleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewHttpWriter(w, r)
		db, err := database.GetDBConn(r.Context())
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

		err = utils.Validator.Struct(&params)
		if err != nil {
			writer.WriteErrorBadRequest(
				err,
				"post /auth/google/sign-in",
				&params,
			)
			return
		}

		profile, err := google.GetPeopleProfile(params.AccessToken)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		socialAccountService := NewSocialAccountService(db)
		socialAccount, err := socialAccountService.FindOneByProviderWithSocialId(
			Key_Provider_GOOGLE,
			profile.SocialId,
		)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		userService := user.NewUserService(db)
		user, err := userService.FindOneById(socialAccount.UserId)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		accessToken, refreshToken, err := token.GenerateTokens(token.GenerateTokenParams{
			UserUUID:    user.UUID,
			DisplayName: user.DisplayName,
			Email:       user.Email,
			PhotoUrl:    utils.NormalizeNullString(user.PhotoUrl),
		})
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
				Email:       user.Email,
				PhotoUrl:    utils.NormalizeNullString(user.PhotoUrl),
				AccessToken: accessToken,
			},
		})
	}
}

func postGoogleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewHttpWriter(w, r)
		db, err := database.GetDBConn(r.Context())
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

		err = utils.Validator.Struct(&params)
		if err != nil {
			writer.WriteErrorBadRequest(
				err,
				"post /auth/google/sign-up",
				&params,
			)
			return
		}

		authService := NewAuthService()
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

		profile, err := google.GetPeopleProfile(params.AccessToken)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-up",
			)
			return
		}

		userService := user.NewUserService(db)
		lastInsertUserId, err := userService.Create(user.User{
			UUID:        utils.GenerateUUIDString(),
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

		user, err := userService.FindOneById(lastInsertUserId)
		if err != nil {
			writer.WriteError(
				err,
				"post /auth/google/sign-up",
			)
			return
		}

		socialAccountService := NewSocialAccountService(db)
		_, err = socialAccountService.Create(SocialAccount{
			UserId:   user.ID,
			UUID:     utils.GenerateUUIDString(),
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

		accessToken, refreshToken, err := token.GenerateTokens(token.GenerateTokenParams{
			UserUUID:    user.UUID,
			DisplayName: user.DisplayName,
			Email:       user.Email,
			PhotoUrl:    utils.NormalizeNullString(user.PhotoUrl),
		})
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
				Email:       user.Email,
				PhotoUrl:    utils.NormalizeNullString(user.PhotoUrl),
				AccessToken: accessToken,
			},
		})
	}
}
