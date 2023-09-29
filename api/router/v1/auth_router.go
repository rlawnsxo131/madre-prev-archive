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
		r.Post("/signup", func(w http.ResponseWriter, r *http.Request) {})
		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {})
	})

	return ar
}

func (ar *authRouter) checkRegistration() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		provider := chi.URLParam(r, "provider")
		var body = struct {
			AccessToken string `json:"accessToken" validate:"min=8,max=50,regexp=^[a-zA-Z0-9]"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if err := validator.Validate(body); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		httpresponse.Json(w, r,
			httpresponse.NewResponse(
				http.StatusOK,
				map[string]any{
					"provider": provider,
					"token":    body.AccessToken,
				},
			),
		)
	}
}
