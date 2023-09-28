package persistence

import (
	"context"
	"database/sql"
	"time"

	"github.com/rlawnsxo131/madre-server/core/logger"
	"github.com/rs/zerolog"
)

type QueryLayer interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type QueryOptions struct {
	DB     QueryLayer
	WithTx bool
}

// query logger
type QueryLogger struct {
	l *zerolog.Logger
}

func NewQueryLogger() *QueryLogger {
	return &QueryLogger{
		l: logger.NewDefaultZeroLogger(),
	}
}

func (ql *QueryLogger) Logging(sql string, args ...any) {
	ql.l.Log().
		Str("time", time.Now().UTC().Format(time.RFC3339Nano)).
		Str("SQL", sql).
		Any("Args", args).
		Send()
}
