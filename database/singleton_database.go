package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

// TODO: Thinking about redefining the database and working on it

type singletonDatabase struct {
	Database *sqlx.DB
	logger   *zerolog.Logger
}

func (sd *singletonDatabase) Begin() {}

func (sd *singletonDatabase) Commit() {}

func (sd *singletonDatabase) Rollback() {}
