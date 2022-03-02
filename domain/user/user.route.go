package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib"
)

func SetupRoute(v1 *mux.Router) {
	userRouter := v1.NewRoute().PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/{uuid}", get()).Methods("GET")
	userRouter.HandleFunc("/{uuid}", put()).Methods("PUT")
}

func get() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		uuid := vars["uuid"]

		db, err := database.GetDBConn(r.Context())
		if err != nil {
			lib.ResponseErrorWriter(rw, err)
			return
		}

		userService := NewUserService(db)
		user, err := userService.FindOneByUUID(uuid)
		if err != nil {
			lib.ResponseErrorWriter(rw, err)
			return
		}

		lib.ResponseJsonCompressWriter(rw, r, user)
	}
}

func put() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {}
}
