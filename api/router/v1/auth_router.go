package routerv1

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/rlawnsxo131/madre-server/api/service/command"
	"github.com/rlawnsxo131/madre-server/core/httpresponse"
	"github.com/rlawnsxo131/madre-server/domain/persistence"
)

type authRouter struct {
	validator          *validator.Validate
	userCommandService *command.UserCommandService
}

func NewAuthRouter(r chi.Router, db persistence.Conn) *authRouter {
	ar := &authRouter{
		validator:          validator.New(validator.WithRequiredStructEnabled()),
		userCommandService: command.NewUserCommandService(db),
	}

	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup-login/{provider}", ar.signupLogin())
		r.Delete("/logout", ar.logout())
		r.Delete("/", ar.deleteAccount())
	})

	return ar
}

func (ar *authRouter) signupLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)

		var body struct {
			AccessToken string `json:"accessToken"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		var params = struct {
			Provider    string `validate:"required,oneof=GOOGLE"`
			AccessToken string `validate:"required,min=8,max=25"`
		}{
			Provider:    chi.URLParam(r, "provider"),
			AccessToken: body.AccessToken,
		}
		if err := ar.validator.Struct(params); err != nil {
			var fields []string
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				for _, validationError := range validationErrors {
					fields = append(fields, strings.ToLower(validationError.Field()))
				}
				log.Printf("fields: %+v", fields)
			}

			rw.ERROR(
				err,
				httpresponse.NewError(
					http.StatusUnprocessableEntity,
					map[string][]string{
						"fields": fields,
					},
					"잘못된 형식입니다",
				),
			)
			return
		}

		rw.JSON(
			httpresponse.New(
				http.StatusOK,
				map[string]any{
					"provider":    params.Provider,
					"accessToken": params.AccessToken,
				},
			),
		)
	}
}

func (ar *authRouter) logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (ar *authRouter) deleteAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
