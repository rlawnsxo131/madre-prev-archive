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
	"github.com/rlawnsxo131/madre-server-v2/lib/syncmap"
)

var sqlxDB *sqlx.DB

func GetDB() (*sqlx.DB, error) {
	if sqlxDB == nil {
		db, err := sqlx.Connect("mysql", "root:1234@/madre?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true")
		if err != nil {
			return nil, errors.Wrap(err, "sqlx: connect fail")
		}

		err = db.Ping()
		if err != nil {
			return nil, errors.Wrap(err, "sqlx: ping fail")
		}

		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		db.SetConnMaxLifetime(time.Minute)
		sqlxDB = db
	}
	return sqlxDB, nil
}

func GetDBConn(ctx context.Context) (*sqlx.DB, error) {
	syncMap, err := syncmap.GetFromHttpContext(ctx)
	if err != nil {
		return nil, err
	}

	if sqlxDB, ok := syncMap.Load(constants.Key_HttpContextDB); ok {
		if sqlxDB, ok := sqlxDB.(*sqlx.DB); ok {
			return sqlxDB, nil
		}
	}
	return nil, errors.New("DB is not exist")
}

func ExcuteInitSQL(db *sqlx.DB) {
	file, err := ioutil.ReadFile("./database/init.sql")
	if err != nil {
		panic(err)
	}

	queries := strings.Split(string(file), "\n\n")
	for _, query := range queries {
		sqlxDB.MustExec(query)
	}
}
