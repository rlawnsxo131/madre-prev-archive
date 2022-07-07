package apiv1

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/external/engine/httpresponse"
	commandentity "github.com/rlawnsxo131/madre-server-v3/internal/application/entity/command"
	commandservice "github.com/rlawnsxo131/madre-server-v3/internal/application/service/command"
	queryservice "github.com/rlawnsxo131/madre-server-v3/internal/application/service/query"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/account"
	"github.com/rlawnsxo131/madre-server-v3/lib/social"
	"github.com/rlawnsxo131/madre-server-v3/lib/token"
	"github.com/rlawnsxo131/madre-server-v3/utils"
)

type authRoute struct {
	accountCommandService account.AccountCommandService
	accountQueryService   account.AccountQueryService
}

func NewAuthRoute(db rdb.Database) *authRoute {
	return &authRoute{
		commandservice.NewAccountCommandService(db),
		queryservice.NewAccountQueryService(db),
	}
}

func (ar *authRoute) Register(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Get("/", ar.Get())
		r.Delete("/", ar.Delete())
		r.Post("/google/check", ar.PostGoogleCheck())
		r.Post("/google/sign-in", ar.PostGoogleSignIn())
		r.Post("/google/sign-up", ar.PostGoogleSignUp())
	})
}

func (ar *authRoute) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)
		p := token.ProfileCtx(r.Context())

		rw.Write(p)
	}
}

func (ar *authRoute) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)
		p := token.ProfileCtx(r.Context())
		if p == nil {
			rw.ErrorUnauthorized(
				errors.New("not found token profile"),
			)
			return
		}
		token.NewManager().ResetCookies(w)

		rw.Write(struct{}{})
	}
}

func (ar *authRoute) PostGoogleCheck() http.HandlerFunc {
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
				errors.Wrap(err, "params validate error"),
			)
			return
		}

		ggp, err := social.NewGooglePeopleAPI(params.AccessToken).Do()
		if err != nil {
			rw.Error(err)
			return
		}

		exist, err := ar.accountQueryService.GetExistsSocialAccountBySocialIdAndProvider(
			ggp.SocialID,
			account.SOCIAL_ACCOUNT_PROVIDER_GOOGLE,
		)
		if err != nil {
			rw.Error(err)
			return
		}

		rw.Write(map[string]bool{
			"exist": exist,
		})
	}
}

func (ar *authRoute) PostGoogleSignIn() http.HandlerFunc {
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
				errors.Wrap(err, "params validate error"),
			)
			return
		}

		ggp, err := social.NewGooglePeopleAPI(params.AccessToken).Do()
		if err != nil {
			rw.Error(err)
			return
		}

		sa, err := ar.accountQueryService.GetSocialAccountBySocialIdAndProvider(
			ggp.SocialID,
			account.SOCIAL_ACCOUNT_PROVIDER_GOOGLE,
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

		u, err := ar.accountQueryService.GetUserById(sa.UserID)
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

		p := token.NewProfile(
			u.ID,
			u.Username,
			utils.NormalizeNullString(u.PhotoUrl),
		)
		tokenManager := token.NewManager()
		err = tokenManager.GenerateAndSetCookies(p, w)
		if err != nil {
			rw.Error(err)
			return
		}

		rw.Write(p)
	}
}

func (ar *authRoute) PostGoogleSignUp() http.HandlerFunc {
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
				errors.Wrap(err, "params validate error"),
			)
			return
		}

		ggp, err := social.NewGooglePeopleAPI(params.AccessToken).Do()
		if err != nil {
			rw.Error(err)
			return
		}

		u := commandentity.NewCreateAccountUser(
			ggp.Email,
			ggp.DisplayName,
			params.Username,
			ggp.PhotoUrl,
		)

		valid, err := u.ValidateUsername()
		if err != nil {
			rw.Error(err)
			return
		}
		if !valid {
			rw.ErrorBadRequest(
				errors.New("username validation error"),
			)
		}

		exist, err := ar.accountQueryService.GetExistsUserByUsername(params.Username)
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

		exist, err = ar.accountQueryService.GetExistsSocialAccountBySocialIdAndProvider(
			ggp.SocialID,
			account.SOCIAL_ACCOUNT_PROVIDER_GOOGLE,
		)
		if err != nil {
			rw.Error(err)
			return
		}
		if exist {
			rw.ErrorUnprocessableEntity(
				errors.Wrap(err, "socialaccount is exist"),
			)
			return
		}

		sa := commandentity.NewCreateAccountSocialAccount(
			ggp.SocialID,
			account.SOCIAL_ACCOUNT_PROVIDER_GOOGLE,
		)

		ac, err := ar.accountCommandService.CreateAccount(u, sa)
		if err != nil {
			rw.Error(err)
			return
		}

		p := token.NewProfile(
			ac.UserID,
			ac.Username,
			ac.PhotoUrl,
		)
		tokenManager := token.NewManager()
		err = tokenManager.GenerateAndSetCookies(p, w)
		if err != nil {
			rw.Error(err)
			return
		}

		rw.Write(p)
	}
}
