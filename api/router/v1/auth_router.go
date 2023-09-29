package routerv1

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server/api/service"
	"github.com/rlawnsxo131/madre-server/core/httpresponse"
	"github.com/rlawnsxo131/madre-server/domain/persistence"
	"gopkg.in/validator.v2"
)

type authRouter struct {
	userService *service.UserService
}

func NewAuthRouter(r chi.Router, db persistence.Conn) *authRouter {
	ar := &authRouter{
		userService: service.NewUserService(db),
	}

	r.Route("/auth", func(r chi.Router) {
		r.Post("/check-registration/{provider}", ar.checkRegistration())
		r.Post("/signup/{provider}", ar.signup())
		r.Post("/login/{provider}", ar.login())
		r.Delete("/logout", ar.logout())
	})

	return ar
}

func (ar *authRouter) checkRegistration() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		provider := chi.URLParam(r, "provider")
		var body = struct {
			AccessToken string `json:"accessToken" validate:"min=8,max=25,regexp=^[a-zA-Z0-9]"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if err := validator.Validate(body); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}

		jsonRes, _ := json.Marshal(
			httpresponse.New(
				http.StatusOK,
				map[string]any{
					"provider":    provider,
					"accessToken": body.AccessToken,
				},
				nil,
			),
		)
		w.Write(jsonRes)
	}
}

func (ar *authRouter) signup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (ar *authRouter) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (ar *authRouter) logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
