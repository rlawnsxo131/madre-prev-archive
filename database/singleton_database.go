package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

type Database interface {
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	PrepareNamedGet(id *string, query string, args interface{}) error
}

type singletonDatabase struct {
	DB     *sqlx.DB
	logger *zerolog.Logger
}

func (sd *singletonDatabase) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	sd.logger.Debug().Msgf("sql: %s,%+v", query, args)
	return sd.DB.Queryx(query, args...)
}

func (sd *singletonDatabase) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	sd.logger.Debug().Msgf("sql: %s,%+v", query, args)
	return sd.DB.QueryRowx(query, args...)
}

func (sd *singletonDatabase) PrepareNamedGet(id *string, query string, args interface{}) error {
	sd.logger.Debug().Msgf("sql: %s,%+v", query, args)
	stmt, err := sd.DB.PrepareNamed(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	return stmt.Get(id, args)
}
