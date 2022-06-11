package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/external/engine/httpresponse"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/auth"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/repository"
	"github.com/rlawnsxo131/madre-server-v3/lib/social"
	"github.com/rlawnsxo131/madre-server-v3/lib/token"
	"github.com/rlawnsxo131/madre-server-v3/utils"
)

type authGoogleHandler struct {
	socialAccountService auth.SocialAccountService
	userService          user.UserService
}

func NewAuthGoogleHandler(db rdb.Database) *authGoogleHandler {
	return &authGoogleHandler{
		auth.NewSocialAccountService(
			repository.NewSocialAccountCommandRepository(db),
			repository.NewSocialAccountQueryRepository(db),
		),
		user.NewUserService(
			repository.NewUserCommandRepository(db),
			repository.NewUserQueryRepository(db),
		),
	}
}

func (h *authGoogleHandler) Register(r chi.Router) {
	r.Route("/auth/google", func(r chi.Router) {
		r.Post("/check", h.PostGoogleCheck())
		r.Post("/sign-in", h.PostGoogleSignIn())
		r.Post("/sign-up", h.PostGoogleSignUp())
	})
}

func (h *authGoogleHandler) PostGoogleCheck() http.HandlerFunc {
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

		ggp, err := social.NewGooglePeopleAPI(params.AccessToken).Do()
		if err != nil {
			rw.Error(err)
			return
		}

		// if no rows in result set err -> { exist: false }
		sa, err := h.socialAccountService.FindOneBySocialIdAndProvider(
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

func (h *authGoogleHandler) PostGoogleSignIn() http.HandlerFunc {
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

		ggp, err := social.NewGooglePeopleAPI(params.AccessToken).Do()
		if err != nil {
			rw.Error(err)
			return
		}

		sa, err := h.socialAccountService.FindOneBySocialIdAndProvider(
			ggp.SocialID,
			auth.Key_SocialAccount_Provider_GOOGLE,
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

func (h *authGoogleHandler) PostGoogleSignUp() http.HandlerFunc {
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

		ggp, err := social.NewGooglePeopleAPI(params.AccessToken).Do()
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

		sa := auth.SocialAccount{
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
