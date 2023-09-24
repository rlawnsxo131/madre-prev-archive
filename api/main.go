package main

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"github.com/rlawnsxo131/madre-server/core/errorz"
	"github.com/rlawnsxo131/madre-server/core/funk"
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

	u, _ := user.NewUserWithoutId("email", "photoUrl")
	u.SetNewSocialAccount("socialId", "GOOGLE")
	jsonRes, _ := json.Marshal(u)
	log.Println(string(jsonRes))
}
