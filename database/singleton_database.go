package database

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

// TODO: Thinking about redefining the database and working on it

type singletonDatabase struct {
	database *sqlx.DB
	logger   *zerolog.Logger
}

func (sd *singletonDatabase) Begin() {}

func (sd *singletonDatabase) Commit() {}

func (sd *singletonDatabase) Rollback() {}

func (sd *singletonDatabase) initDatabase() {
	sd.database.SetMaxIdleConns(5)
	sd.database.SetMaxOpenConns(5)
	sd.database.SetConnMaxLifetime(time.Minute)
}
