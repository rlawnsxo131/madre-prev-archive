package auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/internal/engine/httpresponse"
	"github.com/rlawnsxo131/madre-server-v3/internal/lib/social"
	"github.com/rlawnsxo131/madre-server-v3/internal/lib/token"
	"github.com/rlawnsxo131/madre-server-v3/internal/pkg/user"
	"github.com/rlawnsxo131/madre-server-v3/internal/utils"
)

type httpGoogleHandler struct {
	socialAccountService SocialAccountService
	userService          user.UserService
}

func NewHTTPGoogleHandler(
	socialAccountRepo SocialAccountRepository,
	userRepo user.UserRepository,
) *httpGoogleHandler {
	socialAccountService := NewSocialAccountService(socialAccountRepo)
	userService := user.NewUserService(userRepo)

	return &httpGoogleHandler{
		socialAccountService,
		userService,
	}
}

func (h *httpGoogleHandler) PostGoogleCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)

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
		sa, err := h.socialAccountService.FindOneBySocialIdAndProvider(
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

func (h *httpGoogleHandler) PostGoogleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)

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
				errors.Wrap(err, "PostGoogleSignInRequestDto validate error"),
			)
			return
		}

		ggp, err := social.NewGoogleApi(params.AccessToken).Do()
		if err != nil {
			rw.Error(err)
			return
		}

		sa, err := h.socialAccountService.FindOneBySocialIdAndProvider(
			ggp.SocialID,
			Key_SocialAccount_Provider_GOOGLE,
		)
		if err != nil {
			rw.Error(err)
			return
		}

		u, err := h.userService.FindOneById(sa.UserID)
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

func (h *httpGoogleHandler) PostGoogleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)

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

		sameNameUser, err := h.userService.FindOneByUsername(params.Username)
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

		userId, err := h.userService.Create(&u)
		if err != nil {
			rw.Error(err)
			return
		}

		user, err := h.userService.FindOneById(userId)
		if err != nil {
			rw.Error(err)
			return
		}

		sa := SocialAccount{
			UserID:   user.ID,
			SocialID: ggp.SocialID,
			Provider: "GOOGLE",
		}
		_, err = h.socialAccountService.Create(&sa)
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
