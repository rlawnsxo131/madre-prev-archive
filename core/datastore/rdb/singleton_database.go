package rdb

import (
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	database              *singletonDatabase
	onceSingletonDatabase sync.Once
)

type singletonDatabase struct {
	Conn *pgxpool.Conn
}

func Database(conn *pgxpool.Conn) *singletonDatabase {
	onceSingletonDatabase.Do(func() {
		database = &singletonDatabase{
			Conn: conn,
		}
	})
	return database
}
