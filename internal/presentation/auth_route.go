package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/external/engine/httpresponse"
	"github.com/rlawnsxo131/madre-server-v3/internal/application"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/account"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/command"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/query"
	"github.com/rlawnsxo131/madre-server-v3/lib/social"
	"github.com/rlawnsxo131/madre-server-v3/lib/token"
	"github.com/rlawnsxo131/madre-server-v3/utils"
)

type authRoute struct {
	socialAccountUseCase account.SocialAccountUseCase
	userUseCase          account.UserUseCase
}

func NewAuthRoute(db rdb.Database) *authRoute {
	return &authRoute{
		application.NewSocialAccountUseCase(
			command.NewSocialAccountCommandRepository(db),
			query.NewSocialAccountQueryRepository(db),
		),
		application.NewUserUseCase(
			command.NewUserCommandRepository(db),
			query.NewUserQueryRepository(db),
		),
	}
}

func (h *authRoute) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)
		p := token.UserProfileCtx(r.Context())

		rw.Write(p)
	}
}

func (h *authRoute) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)
		p := token.UserProfileCtx(r.Context())
		if p == nil {
			rw.ErrorUnauthorized(
				errors.New("not found UserProfile"),
			)
			return
		}
		token.NewManager().ResetCookies(w)

		rw.Write(struct{}{})
	}
}

func (h *authRoute) Register(r chi.Router) {
	r.Route("/auth/google", func(r chi.Router) {
		r.Post("/check", h.PostGoogleCheck())
		r.Post("/sign-in", h.PostGoogleSignIn())
		r.Post("/sign-up", h.PostGoogleSignUp())
	})
}

func (h *authRoute) PostGoogleCheck() http.HandlerFunc {
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
				errors.Wrap(err, "PostGoogleCheck params validate error"),
			)
			return
		}

		ggp, err := social.NewGooglePeopleAPI(params.AccessToken).Do()
		if err != nil {
			rw.Error(err)
			return
		}

		// if no rows in result set err -> { exist: false }
		sa, err := h.socialAccountUseCase.FindOneBySocialIdAndProvider(
			ggp.SocialID,
			account.Key_SocialAccount_Provider_GOOGLE,
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

func (h *authRoute) PostGoogleSignIn() http.HandlerFunc {
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
				errors.Wrap(err, "PostGoogleSignIn params validate error"),
			)
			return
		}

		ggp, err := social.NewGooglePeopleAPI(params.AccessToken).Do()
		if err != nil {
			rw.Error(err)
			return
		}

		sa, err := h.socialAccountUseCase.FindOneBySocialIdAndProvider(
			ggp.SocialID,
			account.Key_SocialAccount_Provider_GOOGLE,
		)
		exist, err := sa.IsExist(err)
		if err != nil {
			rw.Error(err)
			return
		}
		if !exist {
			rw.ErrorBadRequest(
				errors.New("not found socialaccount"),
			)
			return
		}

		u, err := h.userUseCase.FindOneById(sa.UserID)
		exist, err = u.IsExist(err)
		if err != nil {
			rw.Error(err)
			return
		}
		if !exist {
			rw.ErrorBadRequest(
				errors.New("not found user"),
			)
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

func (h *authRoute) PostGoogleSignUp() http.HandlerFunc {
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
				errors.Wrap(err, "PostGoogleSignUpRequest params validate error"),
			)
			return
		}

		ggp, err := social.NewGooglePeopleAPI(params.AccessToken).Do()
		if err != nil {
			rw.Error(err)
			return
		}

		// TODO: already exist social account validation process

		u := account.User{
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

		sameNameUser, err := h.userUseCase.FindOneByUsername(params.Username)
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

		userId, err := h.userUseCase.Create(&u)
		if err != nil {
			rw.Error(err)
			return
		}

		user, err := h.userUseCase.FindOneById(userId)
		if err != nil {
			rw.Error(err)
			return
		}

		sa := account.SocialAccount{
			UserID:   user.ID,
			SocialID: ggp.SocialID,
			Provider: account.Key_SocialAccount_Provider_GOOGLE,
		}
		_, err = h.socialAccountUseCase.Create(&sa)
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
