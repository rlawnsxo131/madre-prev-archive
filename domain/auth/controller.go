package auth

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/lib/token"
)

type Controller interface {
	Get() http.HandlerFunc
	Delete() http.HandlerFunc
}

type controller struct {
	db database.Database
}

func NewController(db database.Database) Controller {
	return &controller{
		db: db,
	}
}

func (c *controller) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
		p := token.UserProfileCtx(r.Context())

		rw.Write(p)
	}
}

func (c *controller) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
		p := token.UserProfileCtx(r.Context())
		if p == nil {
			rw.ErrorUnauthorized(
				errors.New("not found userProfile"),
			)
			return
		}
		token.NewManager().ResetCookies(w)

		rw.Write(map[string]interface{}{})
	}
}
