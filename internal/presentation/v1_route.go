package presentation

import (
	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
)

type v1Route struct {
	r  chi.Router
	db rdb.Database
}

func NewV1Route(r chi.Router, db rdb.Database) *v1Route {
	return &v1Route{r, db}
}

func (v1r *v1Route) Register() {
	v1r.r.Route("/v1", func(r chi.Router) {
		NewAuthRoute(v1r.db).Register(r)
	})
}
