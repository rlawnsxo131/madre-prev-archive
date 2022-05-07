package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
)

func ApplyRoutes(v1 *mux.Router) {
	route := v1.NewRoute().PathPrefix("/user").Subrouter()

	route.HandleFunc("/{id}", get()).Methods("GET")
	route.HandleFunc("/{id}", put()).Methods("PUT", "OPTIONS")
}

func get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
		vars := mux.Vars(r)
		id := vars["id"]

		db, err := database.LoadFromHttpCtx(r.Context())
		if err != nil {
			rw.Error(err, "get /user/{id}")
			return
		}

		userUseCase := NewUseCase(db)
		user, err := userUseCase.FindOneById(id)
		if err != nil {
			rw.Error(err, "get /user/{id}")
			return
		}

		rw.Compress(user)
	}
}

func put() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
