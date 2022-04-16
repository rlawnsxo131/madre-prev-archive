package logger

import (
	"database/sql"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type QueryLogger interface {
	NamedExec(query string, args interface{}) (sql.Result, error)
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
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

func (q *queryLogger) NamedExec(query string, args interface{}) (sql.Result, error) {
	q.logger.Debug().Msgf("sql: %s,%+v", query, args)
	return q.db.NamedExec(query, args)
}

func (q *queryLogger) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	q.logger.Debug().Msgf("sql: %s,%+v", query, args)
	return q.db.Queryx(query, args)
}

func (q *queryLogger) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	q.logger.Debug().Msgf("sql: %s,%+v", query, args)
	return q.db.QueryRowx(query, args...)
}
