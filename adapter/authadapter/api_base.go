package authadapter

import (
	"errors"
	"net/http"

	"github.com/rlawnsxo131/madre-server-v3/lib/response"
	"github.com/rlawnsxo131/madre-server-v3/lib/token"
)

type baseHandler struct{}

func NewBaseHandler() *baseHandler {
	return &baseHandler{}
}

func (h *baseHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
		p := token.UserProfileCtx(r.Context())

		rw.Write(p)
	}
}

func (h *baseHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
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
