package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v3/core/server/httpresponse"
)

type healthRoute struct{}

func NewHealthRoute() *healthRoute {
	return &healthRoute{}
}

func (hr *healthRoute) Register(r chi.Router) {
	r.Route("/health", func(r chi.Router) {
		r.Get("/", hr.get())
	})
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

		httpresponse.NewWriter(w, r).Write(
			httpresponse.NewResponse(
				http.StatusOK,
				data,
			),
		)
	}
}
