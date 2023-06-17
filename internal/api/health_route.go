package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v3/core/server/httpresponse"
)

type healthRoute struct{}

func NewHealthRoute(r chi.Router) *healthRoute {
	hr := &healthRoute{}

	r.Route("/health", func(r chi.Router) {
		r.Get("/", hr.get())
	})

	return hr
}

func (hr *healthRoute) get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Proto":   r.Proto,
			"Method":  r.Method,
			"Host":    r.Host,
			"Origin":  r.Header.Get("Origin"),
			"Path":    r.URL.Path,
			"Referer": r.Header.Get("Referer"),
			"Cookies": fmt.Sprint(r.Cookies()),
		}

		// Example
		// id, err := uuid.Parse("asdf")

		// db, _ := rdb.DBInstance()
		// repo := queryrepository.NewUserQueryRepository(db)

		// if u, _ := repo.FindById("959075a4-8de4-4edc-a95c-1a122275762a"); u != nil {
		// 	log.Println(u)
		// }

		httpresponse.NewWriter(w, r).Json(
			httpresponse.New(
				http.StatusOK,
				data,
			),
		)
	}
}
