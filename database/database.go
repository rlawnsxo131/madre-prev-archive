package database

import (
	"context"
	"io/ioutil"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rlawnsxo131/madre-server-v2/constants"
)

var sqlxDb *sqlx.DB

func GetDB() *sqlx.DB {
	if sqlxDb == nil {
		db, err := sqlx.Connect("mysql", "root:1234@/madre?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true")
		if err != nil {
			panic(err)
		}
		err = db.Ping()
		if err != nil {
			panic(err)
		}
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		db.SetConnMaxLifetime(time.Minute)
		sqlxDb = db
	}
	return sqlxDb
}

func GetDBConn(ctx context.Context) *sqlx.DB {
	v := ctx.Value(constants.DBContextKey)
	if v == nil {
		panic("DB is not exist")
	}
	if sqlxDb, ok := v.(*sqlx.DB); ok {
		return sqlxDb
	}
	panic("DB is not exist")
}

func ExcuteInitSQL(db *sqlx.DB) {
	file, err := ioutil.ReadFile("./database/init.sql")
	if err != nil {
		panic(err)
	}
	queries := strings.Split(string(file), "\n\n")
	for _, query := range queries {
		sqlxDb.MustExec(query)
	}
}
