package apiv1

import (
	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
)

type apiv1 struct {
	r  chi.Router
	db rdb.Database
}

func NewAPI(r chi.Router, db rdb.Database) *apiv1 {
	return &apiv1{r, db}
}

func (v1r *apiv1) Register() {
	v1r.r.Route("/v1", func(r chi.Router) {
		NewAuthRoute(v1r.db).Register(r)
	})
}
