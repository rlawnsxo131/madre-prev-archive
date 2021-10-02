package main

import (
	"fmt"
	"net/http"

	"github.com/rlawnsxo131/madre-server/src/graphql"
)

func main() {
	h, err := graphql.NewHandler()

	if err != nil {
		fmt.Println("err: ", err)
	}

	http.Handle("/graphql", h)
	http.ListenAndServe(":3001", nil)
}
