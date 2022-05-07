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
		writer := response.NewWriter(w, r)
		userProfile := token.LoadCtxUserProfile(r.Context())

		writer.Compress(
			map[string]interface{}{
				"user_token_profile": userProfile,
			},
		)
	}
}

func delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewWriter(w, r)
		userProfile := token.LoadCtxUserProfile(r.Context())

		if userProfile == nil {
			writer.ErrorUnauthorized(
				errors.New("not found userTokenProfile"),
				"delete /auth",
				userProfile,
			)
			return
		}

		token.ResetTokenCookies(w)
		writer.Compress(map[string]interface{}{})
	}
}

func postGoogleCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewWriter(w, r)
		db, err := database.LoadFromHttpCtx(r.Context())
		if err != nil {
			writer.Error(
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
			writer.Error(
				errors.WithStack(err),
				"post /auth/google/check",
				"decode params error",
			)
			return
		}

		err = utils.NewValidator().Struct(&params)
		if err != nil {
			writer.ErrorBadRequest(
				err,
				"post /auth/google/check",
				params,
			)
			return
		}

		googleProfile, err := social.NewGoogleApi().Do(params.AccessToken)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/check",
			)
			return
		}

		// if no rows in result set err -> { exist: false }
		socialUseCase := socialaccount.NewUseCase(db)
		sa, err := socialUseCase.FindOneByProviderWithSocialId(
			socialaccount.Key_Provider_GOOGLE,
			googleProfile.SocialId,
		)
		exist, err := sa.IsExist(err)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/check",
			)
			return
		}

		writer.Compress(map[string]bool{
			"exist": exist,
		})
	}
}

func postGoogleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewWriter(w, r)
		db, err := database.LoadFromHttpCtx(r.Context())
		if err != nil {
			writer.Error(
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
			writer.Error(
				errors.WithStack(err),
				"post /auth/google/sign-in",
				"decode params error",
			)
			return
		}

		err = utils.NewValidator().Struct(&params)
		if err != nil {
			writer.ErrorBadRequest(
				err,
				"post /auth/google/sign-in",
				params,
			)
			return
		}

		googleProfile, err := social.NewGoogleApi().Do(params.AccessToken)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		socialUseCase := socialaccount.NewUseCase(db)
		sa, err := socialUseCase.FindOneByProviderWithSocialId(
			socialaccount.Key_Provider_GOOGLE,
			googleProfile.SocialId,
		)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		userUserCase := user.NewUseCase(db)
		u, err := userUserCase.FindOneById(sa.UserID)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		userProfile := token.UserProfile{
			UserID:      u.ID,
			DisplayName: u.DisplayName,
			PhotoUrl:    utils.NormalizeNullString(u.PhotoUrl),
		}
		accessToken, refreshToken, err := token.GenerateTokens(&userProfile)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		token.SetTokenCookies(w, accessToken, refreshToken)

		writer.Compress(map[string]interface{}{
			"user_token_profile": userProfile,
		})
	}
}

func postGoogleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewWriter(w, r)
		db, err := database.LoadFromHttpCtx(r.Context())
		if err != nil {
			writer.Error(
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
			writer.Error(
				err,
				"post /auth/google/sign-up",
				"decode params error",
			)
			return
		}

		err = utils.NewValidator().Struct(&params)
		if err != nil {
			writer.ErrorBadRequest(
				err,
				"post /auth/google/sign-up",
				params,
			)
			return
		}

		googleProfile, err := social.NewGoogleApi().Do(params.AccessToken)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-up",
			)
			return
		}

		u := &user.User{
			Email:       googleProfile.Email,
			OriginName:  utils.NewNullString(googleProfile.DisplayName),
			DisplayName: params.DisplayName,
			PhotoUrl:    utils.NewNullString(googleProfile.PhotoUrl),
		}
		valid, err := u.ValidateDisplayName()
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-up",
				"username validate error",
			)
			return
		}
		if !valid {
			writer.ErrorBadRequest(
				errors.New("username validation error"),
				"post /auth/google/sign-up",
				params,
			)
			return
		}

		userUseCase := user.NewUseCase(db)
		userId, err := userUseCase.Create(u)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-up",
			)
			return
		}

		user, err := userUseCase.FindOneById(userId)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-up",
			)
			return
		}

		sa := socialaccount.SocialAccount{
			UserID:   user.ID,
			Provider: "GOOGLE",
			SocialId: googleProfile.SocialId,
		}
		socialUseCase := socialaccount.NewUseCase(db)
		_, err = socialUseCase.Create(&sa)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-up",
			)
			return
		}

		userProfile := token.UserProfile{
			UserID:      user.ID,
			DisplayName: user.DisplayName,
			PhotoUrl:    utils.NormalizeNullString(user.PhotoUrl),
		}
		accessToken, refreshToken, err := token.GenerateTokens(&userProfile)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		token.SetTokenCookies(w, accessToken, refreshToken)

		writer.Compress(map[string]interface{}{
			"user_token_profile": userProfile,
		})
	}
}
