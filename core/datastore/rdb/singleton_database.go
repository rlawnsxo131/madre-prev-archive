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
	pool *pgxpool.Pool
}

func Database(p *pgxpool.Pool) *singletonDatabase {
	onceSingletonDatabase.Do(func() {
		database = &singletonDatabase{
			pool: p,
		}
	})
	return database
}
