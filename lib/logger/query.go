package logger

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type QueryLogger interface {
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	PrepareNamedGet(id *string, query string, args interface{}) error
}

type queryLogger struct {
	db     *sqlx.DB
	logger *zerolog.Logger
}

func NewQueryLogger(db *sqlx.DB) QueryLogger {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &queryLogger{
		db:     db,
		logger: &logger,
	}
}

func (q *queryLogger) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	q.logger.Debug().Msgf("sql: %s,%+v", query, args)
	return q.db.Queryx(query, args...)
}

func (q *queryLogger) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	q.logger.Debug().Msgf("sql: %s,%+v", query, args)
	return q.db.QueryRowx(query, args...)
}

func (q *queryLogger) PrepareNamedGet(id *string, query string, args interface{}) error {
	q.logger.Debug().Msgf("sql: %s,%+v", query, args)
	stmt, err := q.db.PrepareNamed(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	return stmt.Get(id, args)
}
