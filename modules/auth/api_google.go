package auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/lib/social"
	"github.com/rlawnsxo131/madre-server-v2/lib/token"
	"github.com/rlawnsxo131/madre-server-v2/modules/user"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type (
	PostGoogleCheckRequestParams struct {
		AccessToken string `json:"access_token" validate:"required,min=50"`
	}
	PostGoogleSignInRequestParams struct {
		AccessToken string `json:"access_token" validate:"required,min=50"`
	}
	PostGoogleSignUpRequestParams struct {
		AccessToken string `json:"access_token" validate:"required,min=50"`
		Username    string `json:"username" validate:"required,max=20,min=1"`
	}
)

type googleHandler struct {
	db database.Database
}

func NewGoogleHandler(db database.Database) *googleHandler {
	return &googleHandler{
		db: db,
	}
}

func (h *googleHandler) PostGoogleCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)

		var params PostGoogleCheckRequestParams
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
				errors.Wrap(err, "PostGoogleCheckRequestParams validate error"),
			)
			return
		}

		ggp, err := social.NewGoogleApi(params.AccessToken).Do()
		if err != nil {
			rw.Error(err)
			return
		}

		// if no rows in result set err -> { exist: false }
		socialUseCase := NewSocialAccountUseCase(
			NewSocialAccountRepository(h.db),
		)
		sa, err := socialUseCase.FindOneBySocialIdAndProvider(
			ggp.SocialID,
			Key_SocialAccount_Provider_GOOGLE,
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

func (h *googleHandler) PostGoogleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)

		var params PostGoogleSignInRequestParams
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
				errors.Wrap(err, "PostGoogleSignInRequestParams validate error"),
			)
			return
		}

		ggp, err := social.NewGoogleApi(params.AccessToken).Do()
		if err != nil {
			rw.Error(err)
			return
		}

		socialUseCase := NewSocialAccountUseCase(
			NewSocialAccountRepository(h.db),
		)
		sa, err := socialUseCase.FindOneBySocialIdAndProvider(
			ggp.SocialID,
			Key_SocialAccount_Provider_GOOGLE,
		)
		if err != nil {
			rw.Error(err)
			return
		}

		userUseCase := user.NewUserUseCase(
			user.NewUserRepository(h.db),
		)
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
		tokenManager := token.NewManager()
		err = tokenManager.GenerateAndSetCookies(&p, w)
		if err != nil {
			rw.Error(err)
			return
		}

		rw.Write(&p)
	}
}

func (h *googleHandler) PostGoogleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)

		var params PostGoogleSignUpRequestParams
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
				errors.Wrap(err, "PostGoogleSignUpRequestParams validate error"),
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

		userUseCase := user.NewUserUseCase(
			user.NewUserRepository(h.db),
		)
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

		sa := SocialAccount{
			UserID:   user.ID,
			SocialID: ggp.SocialID,
			Provider: "GOOGLE",
		}
		socialUseCase := NewSocialAccountUseCase(
			NewSocialAccountRepository(h.db),
		)
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
		tokenManager := token.NewManager()
		err = tokenManager.GenerateAndSetCookies(&p, w)
		if err != nil {
			rw.Error(err)
			return
		}

		rw.Write(&p)
	}
}
