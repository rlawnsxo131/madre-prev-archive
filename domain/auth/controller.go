package auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/domain/user"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/lib/social"
	"github.com/rlawnsxo131/madre-server-v2/lib/token"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type (
	PostGoogleCheckRequestDto struct {
		AccessToken string `json:"access_token" validate:"required,min=50"`
	}
	PostGoogleSignInRequestDto struct {
		AccessToken string `json:"access_token" validate:"required,min=50"`
	}
	PostGoogleSignUpRequestDto struct {
		AccessToken string `json:"access_token" validate:"required,min=50"`
		Username    string `json:"username" validate:"required,max=20,min=1"`
	}
)

type controller struct {
	db database.Database
}

func NewController(db database.Database) *controller {
	return &controller{
		db: db,
	}
}

func (c *controller) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
		p := token.UserProfileCtx(r.Context())

		rw.Write(p)
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
		token.NewManager().ResetCookies(w)

		rw.Write(struct{}{})
	}
}

func (c *controller) PostGoogleCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)

		var params PostGoogleCheckRequestDto
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
				errors.Wrap(err, "PostGoogleCheckRequestDto validate error"),
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
			NewSocialAccountRepository(c.db),
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

func (c *controller) PostGoogleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)

		var params PostGoogleSignInRequestDto
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
				errors.Wrap(err, "PostGoogleSignInRequestDto validate error"),
			)
			return
		}

		ggp, err := social.NewGoogleApi(params.AccessToken).Do()
		if err != nil {
			rw.Error(err)
			return
		}

		socialUseCase := NewSocialAccountUseCase(
			NewSocialAccountRepository(c.db),
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
			user.NewUserRepository(c.db),
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

func (c *controller) PostGoogleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)

		var params PostGoogleSignUpRequestDto
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
				errors.Wrap(err, "PostGoogleSignUpRequestDto validate error"),
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
			user.NewUserRepository(c.db),
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
			NewSocialAccountRepository(c.db),
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
