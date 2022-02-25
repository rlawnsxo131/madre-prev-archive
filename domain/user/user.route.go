package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib"
)

func SetupRoute(v1 *mux.Router) {
	userRouter := v1.NewRoute().PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/{id}", getOne()).Methods("GET")
}

func getOne() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		db, err := database.GetDBConn(r.Context())
		if err != nil {
			lib.ResponseErrorWriter(rw, err)
			return
		}

		userService := NewUserService(db)
		user, err := userService.FindOne(id)
		if err != nil {
			lib.ResponseErrorWriter(rw, err)
			return
		}

		lib.ResponseJsonCompressWriter(rw, r, user)
	}
}
