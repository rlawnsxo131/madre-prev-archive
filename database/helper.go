package database

import (
	"io/ioutil"
	"strings"

	"github.com/jmoiron/sqlx"
)

func ExcuteInitSQL(db *sqlx.DB) {
	file, err := ioutil.ReadFile("./database/init.sql")
	if err != nil {
		panic(err)
	}

	queries := strings.Split(string(file), "\n\n")
	for _, query := range queries {
		db.MustExec(query)
	}
}
