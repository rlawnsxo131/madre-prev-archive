package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
)

func ApplyRoutes(v1 *mux.Router) {
	userRoute := v1.NewRoute().PathPrefix("/user").Subrouter()

	userRoute.HandleFunc("/{uuid}", get()).Methods("GET")
	userRoute.HandleFunc("/{uuid}", put()).Methods("PUT", "OPTIONS")
}

func get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewHttpWriter(w, r)
		vars := mux.Vars(r)
		uuid := vars["uuid"]

		db, err := database.GetDBConn(r.Context())
		if err != nil {
			writer.WriteError(err, "get /user/{uuid}")
			return
		}

		userService := NewUserService(db)
		user, err := userService.FindOneByUUID(uuid)
		if err != nil {
			writer.WriteError(err, "get /user/{uuid}")
			return
		}

		writer.WriteCompress(user)
	}
}

func put() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
