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
		uTokenProfile := token.LoadCtxUserTokenProfile(r.Context())

		writer.Compress(
			map[string]interface{}{
				"user_token_profile": uTokenProfile,
			},
		)
	}
}

func delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewWriter(w, r)
		uTokenProfile := token.LoadCtxUserTokenProfile(r.Context())

		if uTokenProfile == nil {
			writer.ErrorUnauthorized(
				errors.New("not found userTokenProfile"),
				"delete /auth",
				uTokenProfile,
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

		gProfile, err := social.NewGoogleApi().Do(params.AccessToken)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/check",
			)
			return
		}

		// if no rows in result set err -> { exist: false }
		saService := socialaccount.NewService(db)
		sa, err := saService.FindOneByProviderWithSocialId(
			socialaccount.Key_Provider_GOOGLE,
			gProfile.SocialId,
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

		gProfile, err := social.NewGoogleApi().Do(params.AccessToken)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		saService := socialaccount.NewService(db)
		sa, err := saService.FindOneByProviderWithSocialId(
			socialaccount.Key_Provider_GOOGLE,
			gProfile.SocialId,
		)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		uService := user.NewService(db)
		u, err := uService.FindOneById(sa.UserID)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		uTokenProfile := token.UserTokenProfile{
			UserID:      u.ID,
			DisplayName: u.DisplayName,
			PhotoUrl:    utils.NormalizeNullString(u.PhotoUrl),
		}
		accessToken, refreshToken, err := token.GenerateTokens(&uTokenProfile)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		token.SetTokenCookies(w, accessToken, refreshToken)

		writer.Compress(map[string]interface{}{
			"user_token_profile": uTokenProfile,
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

		gProfile, err := social.NewGoogleApi().Do(params.AccessToken)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-up",
			)
			return
		}

		u := &user.User{
			Email:       gProfile.Email,
			OriginName:  utils.NewNullString(gProfile.DisplayName),
			DisplayName: params.DisplayName,
			PhotoUrl:    utils.NewNullString(gProfile.PhotoUrl),
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

		uService := user.NewService(db)
		userId, err := uService.Create(u)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-up",
			)
			return
		}

		user, err := uService.FindOneById(userId)
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
			SocialId: gProfile.SocialId,
		}
		saService := socialaccount.NewService(db)
		_, err = saService.Create(&sa)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-up",
			)
			return
		}

		uTokenProfile := token.UserTokenProfile{
			UserID:      user.ID,
			DisplayName: user.DisplayName,
			PhotoUrl:    utils.NormalizeNullString(user.PhotoUrl),
		}
		accessToken, refreshToken, err := token.GenerateTokens(&uTokenProfile)
		if err != nil {
			writer.Error(
				err,
				"post /auth/google/sign-in",
			)
			return
		}

		token.SetTokenCookies(w, accessToken, refreshToken)

		writer.Compress(map[string]interface{}{
			"user_token_profile": uTokenProfile,
		})
	}
}
