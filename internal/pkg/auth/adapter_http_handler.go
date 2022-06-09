package auth

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/internal/engine/httpresponse"
	"github.com/rlawnsxo131/madre-server-v3/internal/lib/token"
)

type httpHandler struct{}

func NewHTTPHandler() *httpHandler {
	return &httpHandler{}
}

func (h *httpHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)
		p := token.UserProfileCtx(r.Context())

		rw.Write(p)
	}
}

func (h *httpHandler) Delete() http.HandlerFunc {
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
