package auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/domain/auth/socialaccount"
	"github.com/rlawnsxo131/madre-server-v2/domain/user"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/lib/social"
	"github.com/rlawnsxo131/madre-server-v2/lib/token"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type GoogleController interface {
	PostGoogleCheck() http.HandlerFunc
	PostGoogleSignIn() http.HandlerFunc
	PostGoogleSignUp() http.HandlerFunc
}

type googleController struct {
	db database.Database
}

func NewGoogleController(db database.Database) GoogleController {
	return &googleController{
		db: db,
	}
}

func (c *googleController) PostGoogleCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)

		var params struct {
			AccessToken string `json:"access_token" validate:"required,min=50"`
		}
		err := json.NewDecoder(r.Body).Decode(&params)
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
		socialReadUseCase := socialaccount.NewReadUseCase(c.db)
		sa, err := socialReadUseCase.FindOneBySocialIdAndProvider(
			ggp.SocialId,
			socialaccount.Key_Provider_GOOGLE,
		)
		exist, err := sa.IsExist(err)
		if err != nil {
			rw.Error(err)
			return
		}

		rw.Write(map[string]bool{
			"exist": exist,
		})
	}
}

func (c *googleController) PostGoogleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)

		var params struct {
			AccessToken string `json:"access_token" validate:"required,min=50"`
		}
		err := json.NewDecoder(r.Body).Decode(&params)
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

		socialReadUseCase := socialaccount.NewReadUseCase(c.db)
		sa, err := socialReadUseCase.FindOneBySocialIdAndProvider(
			ggp.SocialId,
			socialaccount.Key_Provider_GOOGLE,
		)
		if err != nil {
			rw.Error(err)
			return
		}

		userReadUseCase := user.NewReadUseCase(c.db)
		u, err := userReadUseCase.FindOneById(sa.UserID)
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

		rw.Write(map[string]interface{}{
			"user_profile": &p,
		})
	}
}

func (c *googleController) PostGoogleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)

		var params struct {
			AccessToken string `json:"access_token" validate:"required,min=50"`
			Username    string `json:"username" validate:"required,max=20,min=1"`
		}
		err := json.NewDecoder(r.Body).Decode(&params)
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

		userReadUseCase := user.NewReadUseCase(c.db)
		sameNameUser, err := userReadUseCase.FindOneByUsername(params.Username)
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

		userWriteUseCase := user.NewWriteUseCase(c.db)
		userId, err := userWriteUseCase.Create(&u)
		if err != nil {
			rw.Error(err)
			return
		}

		user, err := userReadUseCase.FindOneById(userId)
		if err != nil {
			rw.Error(err)
			return
		}

		sa := socialaccount.SocialAccount{
			UserID:   user.ID,
			SocialId: ggp.SocialId,
			Provider: "GOOGLE",
		}
		socialWriteUseCase := socialaccount.NewWriteUseCase(c.db)
		_, err = socialWriteUseCase.Create(&sa)
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

		rw.Write(map[string]interface{}{
			"user_profile": &p,
		})
	}
}
