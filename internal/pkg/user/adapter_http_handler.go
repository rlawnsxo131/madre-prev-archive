package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v3/internal/engine/httpresponse"
)

type httpHandler struct {
	userService UserService
}

func NewHTTPHandler(userRepo UserRepository) *httpHandler {
	return &httpHandler{
		NewUserService(userRepo),
	}
}

func (h *httpHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)
		id := chi.URLParam(r, "id")

		u, err := h.userService.FindOneById(id)
		if err != nil {
			rw.Error(err)
			return
		}

		rw.Write(u)

	}
}
