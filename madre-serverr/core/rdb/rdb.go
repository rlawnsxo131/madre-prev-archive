package rdb

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

func CreateConnection(cfg *mysql.Config) (*sql.DB, error) {
	connector, err := mysql.NewConnector(cfg)
	if err != nil {
		return nil, err
	}

	db := sql.OpenDB(connector)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}
