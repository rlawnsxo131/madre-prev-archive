package rdb

import (
	"io/ioutil"
	"strings"

	"github.com/jmoiron/sqlx"
)

func ExcuteInitSQL(db *sqlx.DB) {
	file, err := ioutil.ReadFile("./external/datastore/rdb/init.sql")
	if err != nil {
		panic(err)
	}

	queries := strings.Split(string(file), "\n\n")
	for _, query := range queries {
		db.MustExec(query)
	}
}
