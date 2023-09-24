package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/rlawnsxo131/madre-server/core/errorz"
	"github.com/rlawnsxo131/madre-server/core/funk"
	"github.com/rlawnsxo131/madre-server/core/httpserver"
	"github.com/rlawnsxo131/madre-server/domain/user"
)

func main() {
	log.Println("hello api")

	log.Println(errors.Join(errors.New("11"), errors.New("22")))
	log.Println(errorz.New(errors.New("asdf")))

	ss := funk.Map[int, string]([]int{1, 2, 3, 4}, func(v, idx int, ss []string) string {
		return strconv.Itoa(v)
	})
	log.Println("ss: ", ss)

	u, err := user.NewUserWithoutId("asdf@gmail.com", "photoUrl")
	if err != nil {
		log.Fatal(err)
	}
	u.SetNewSocialAccount("socialId", "GOOGLE")
	jsonRes, _ := json.Marshal(u)
	log.Println(string(jsonRes))

	s := httpserver.New("0.0.0.0:5001")
	s.Route().Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	s.Start()
}
