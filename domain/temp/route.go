package temp

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
)

func ApplyRoutes(v1 *mux.Router) {
	temp := v1.NewRoute().PathPrefix("/temp").Subrouter()

	temp.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewWriter(w, r)
		data := map[string]string{
			"data": "data",
		}
		writer.Compress(data)
	}).Methods("GET")

	temp.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewWriter(w, r)
		data := map[string]string{
			"data": "data",
		}
		writer.Compress(data)
	}).Methods("POST", "OPTIONS")
}
