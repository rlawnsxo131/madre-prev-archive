package logger

import (
	"database/sql"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type QueryLogger interface {
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	NamedExec(query string, args ...interface{}) (sql.Result, error)
}

type queryLogger struct {
	db      *sqlx.DB
	queryer sqlx.Queryer
	logger  zerolog.Logger
}

func NewQueryLogger(db *sqlx.DB) QueryLogger {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &queryLogger{
		db:      db,
		queryer: db,
		logger:  logger,
	}
}

func (q *queryLogger) NamedExec(query string, args ...interface{}) (sql.Result, error) {
	q.logger.Debug().Msgf("sql: %s,%+v", query, args)
	return q.db.NamedExec(query, args)
}

func (q *queryLogger) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	q.logger.Debug().Msgf("sql: %s,%+v", query, args)
	return q.queryer.QueryRowx(query, args...)
}
