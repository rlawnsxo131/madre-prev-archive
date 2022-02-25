package database

import (
	"context"
	"io/ioutil"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/constants"
)

var sqlxDb *sqlx.DB

var (
	ErrdbIsNotExist = errors.New("DB is not exist")
)

func GetDB() (db *sqlx.DB, err error) {
	if sqlxDb == nil {
		db, err := sqlx.Connect("mysql", "root:1234@/madre?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true")
		if err != nil {
			return nil, errors.Wrap(err, "[sqlx connect fail]")
		}

		err = db.Ping()
		if err != nil {
			return nil, errors.Wrap(err, "[db ping fail]")
		}
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		db.SetConnMaxLifetime(time.Minute)
		sqlxDb = db
	}
	return sqlxDb, nil
}

func GetDBConn(ctx context.Context) (db *sqlx.DB, err error) {
	v := ctx.Value(constants.DBContextKey)
	if v == nil {
		return nil, ErrdbIsNotExist
	}
	if sqlxDb, ok := v.(*sqlx.DB); ok {
		return sqlxDb, nil
	}
	return nil, ErrdbIsNotExist
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
