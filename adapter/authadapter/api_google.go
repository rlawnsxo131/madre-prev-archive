package authadapter

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"

	"github.com/rlawnsxo131/madre-server-v3/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/domain/auth"
	"github.com/rlawnsxo131/madre-server-v3/domain/auth/authrepository"
	"github.com/rlawnsxo131/madre-server-v3/domain/user/userrepository"

	"github.com/rlawnsxo131/madre-server-v3/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/lib/response"
	"github.com/rlawnsxo131/madre-server-v3/lib/social"
	"github.com/rlawnsxo131/madre-server-v3/lib/token"
	"github.com/rlawnsxo131/madre-server-v3/utils"
)

type googleHandler struct {
	db rdb.Database
}

func NewGoogleHandler(db rdb.Database) *googleHandler {
	return &googleHandler{
		db: db,
	}
}

func (h *googleHandler) PostGoogleCheck() http.HandlerFunc {
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
		socialUseCase := auth.NewSocialAccountUseCase(
			authrepository.NewSocialAccountRepository(h.db),
		)
		sa, err := socialUseCase.FindOneBySocialIdAndProvider(
			ggp.SocialID,
			auth.Key_SocialAccount_Provider_GOOGLE,
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

		socialUseCase := auth.NewSocialAccountUseCase(
			authrepository.NewSocialAccountRepository(h.db),
		)
		sa, err := socialUseCase.FindOneBySocialIdAndProvider(
			ggp.SocialID,
			auth.Key_SocialAccount_Provider_GOOGLE,
		)
		if err != nil {
			rw.Error(err)
			return
		}

		userUseCase := user.NewUserUseCase(
			userrepository.NewUserRepository(h.db),
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
			userrepository.NewUserRepository(h.db),
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

		sa := auth.SocialAccount{
			UserID:   user.ID,
			SocialID: ggp.SocialID,
			Provider: "GOOGLE",
		}
		socialUseCase := auth.NewSocialAccountUseCase(
			authrepository.NewSocialAccountRepository(h.db),
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
