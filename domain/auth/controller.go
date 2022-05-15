package auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/domain/auth/socialaccount"
	"github.com/rlawnsxo131/madre-server-v2/domain/user"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/lib/social"
	"github.com/rlawnsxo131/madre-server-v2/lib/token"
	"github.com/rlawnsxo131/madre-server-v2/utils"
	"github.com/rs/zerolog"
)

type Controller interface {
	Get() http.HandlerFunc
	Delete() http.HandlerFunc
	PostGoogleCheck() http.HandlerFunc
	PostGoogleSignIn() http.HandlerFunc
	PostGoogleSignUp() http.HandlerFunc
}

type controller struct {
	db database.Database
}

func NewController(db database.Database) Controller {
	return &controller{
		db: db,
	}
}

func (c *controller) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
		p := token.UserProfileCtx(r.Context())

		rw.Compress(
			map[string]interface{}{
				"user_profile": p,
			},
		)
	}
}

func (c *controller) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
		p := token.UserProfileCtx(r.Context())
		if p == nil {
			rw.ErrorUnauthorized(
				errors.New("not found userProfile"),
			)
			return
		}
		token.ResetTokenCookies(w)

		rw.Compress(map[string]interface{}{})
	}
}

func (c *controller) PostGoogleCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)

		var params struct {
			AccessToken string `json:"access_token" validate:"required,min=50"`
		}
		err := json.NewDecoder(r.Body).Decode(&params)
		logger.HTTPLoggerCtx(r.Context()).Add(func(e *zerolog.Event) {
			e.Interface("body", params)
		})
		if err != nil {
			rw.Error(
				errors.Wrap(err, "decode params error"),
			)
			return
		}

		err = validator.New().Struct(&params)
		if err != nil {
			rw.ErrorBadRequest(
				errors.Wrap(err, "access_token validate error"),
			)
			return
		}

		ggp, err := social.NewGoogleApi(params.AccessToken).Do()
		if err != nil {
			rw.Error(err)
			return
		}

		// if no rows in result set err -> { exist: false }
		socialUseCase := socialaccount.NewUseCase(c.db)
		sa, err := socialUseCase.FindOneByProviderWithSocialId(
			socialaccount.Key_Provider_GOOGLE,
			ggp.SocialId,
		)
		exist, err := sa.IsExist(err)
		if err != nil {
			rw.Error(err)
			return
		}

		rw.Compress(map[string]bool{
			"exist": exist,
		})
	}
}

func (c *controller) PostGoogleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)

		var params struct {
			AccessToken string `json:"access_token" validate:"required,min=50"`
		}
		err := json.NewDecoder(r.Body).Decode(&params)
		logger.HTTPLoggerCtx(r.Context()).Add(func(e *zerolog.Event) {
			e.Interface("body", params)
		})
		if err != nil {
			rw.Error(
				errors.Wrap(err, "decode params error"),
			)
			return
		}

		err = validator.New().Struct(&params)
		if err != nil {
			rw.ErrorBadRequest(
				errors.Wrap(err, "access_token validate error"),
			)
			return
		}

		ggp, err := social.NewGoogleApi(params.AccessToken).Do()
		if err != nil {
			rw.Error(err)
			return
		}

		socialUseCase := socialaccount.NewUseCase(c.db)
		sa, err := socialUseCase.FindOneByProviderWithSocialId(
			socialaccount.Key_Provider_GOOGLE,
			ggp.SocialId,
		)
		if err != nil {
			rw.Error(err)
			return
		}

		userUseCase := user.NewUseCase(c.db)
		u, err := userUseCase.FindOneById(sa.UserID)
		if err != nil {
			rw.Error(err)
			return
		}

		p := token.UserProfile{
			UserID:   u.ID,
			Username: u.Username,
			PhotoUrl: utils.NormalizeNullString(u.PhotoUrl),
		}
		actk, rftk, err := token.GenerateTokens(&p)
		if err != nil {
			rw.Error(err)
			return
		}
		token.SetTokenCookies(w, actk, rftk)

		rw.Compress(map[string]interface{}{
			"user_profile": &p,
		})
	}
}

func (c *controller) PostGoogleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)

		var params struct {
			AccessToken string `json:"access_token" validate:"required,min=50"`
			Username    string `json:"username" validate:"required,max=20,min=1"`
		}
		err := json.NewDecoder(r.Body).Decode(&params)
		logger.HTTPLoggerCtx(r.Context()).Add(func(e *zerolog.Event) {
			e.Interface("body", params)
		})
		if err != nil {
			rw.Error(
				errors.Wrap(err, "decode params error"),
			)
			return
		}

		err = validator.New().Struct(&params)
		if err != nil {
			rw.ErrorBadRequest(
				errors.Wrap(err, "access_token, username validate error"),
			)
			return
		}

		ggp, err := social.NewGoogleApi(params.AccessToken).Do()
		if err != nil {
			rw.Error(err)
			return
		}

		u := user.User{
			Email:      ggp.Email,
			OriginName: utils.NewNullString(ggp.DisplayName),
			Username:   params.Username,
			PhotoUrl:   utils.NewNullString(ggp.PhotoUrl),
		}
		valid, err := u.ValidateUsername()
		if err != nil {
			rw.Error(err)
			return
		}
		if !valid {
			rw.ErrorBadRequest(
				errors.New("username validation error"),
			)
			return
		}

		userUseCase := user.NewUseCase(c.db)
		sameNameUser, err := userUseCase.FindOneByUsername(params.Username)
		exist, err := sameNameUser.IsExist(err)
		if err != nil {
			rw.Error(err)
			return
		}
		if exist {
			rw.ErrorConflict(
				errors.Wrap(err, "username is exist"),
			)
			return
		}

		userId, err := userUseCase.Create(&u)
		if err != nil {
			rw.Error(err)
			return
		}

		user, err := userUseCase.FindOneById(userId)
		if err != nil {
			rw.Error(err)
			return
		}

		sa := socialaccount.SocialAccount{
			UserID:   user.ID,
			Provider: "GOOGLE",
			SocialId: ggp.SocialId,
		}
		socialUseCase := socialaccount.NewUseCase(c.db)
		_, err = socialUseCase.Create(&sa)
		if err != nil {
			rw.Error(err)
			return
		}

		p := token.UserProfile{
			UserID:   user.ID,
			Username: user.Username,
			PhotoUrl: utils.NormalizeNullString(user.PhotoUrl),
		}
		actk, rftk, err := token.GenerateTokens(&p)
		if err != nil {
			rw.Error(err)
			return
		}
		token.SetTokenCookies(w, actk, rftk)

		rw.Compress(map[string]interface{}{
			"user_profile": &p,
		})
	}
}
