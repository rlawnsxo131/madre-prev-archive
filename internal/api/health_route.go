package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v3/pkg/core/server/httpresponse"
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

		// conn, _ := db.Conn()
		// tx, _ := conn.BeginTx(context.Background(), pgx.TxOptions{})
		// sql := "INSERT INTO public.user (email, username, photo_url)" +
		// 	" VALUES (@email, @username, @photo_url)" +
		// 	"RETURNING id"

		// var id string

		// row := tx.QueryRow(context.Background(), sql,
		// 	pgx.NamedArgs{
		// 		"email":     "email2",
		// 		"username":  "username2",
		// 		"photo_url": "photo_url",
		// 	})

		// if err := row.Scan(&id); err != nil {
		// 	log.Println(err)
		// } else {
		// 	log.Println(id)
		// }

		httpresponse.NewWriter(w, r).Json(
			httpresponse.NewResponse(
				http.StatusOK,
				data,
			),
		)
	}
}
