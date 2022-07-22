package rdb

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

const (
	KEY_DATABASE_CONN_CTX = "KEY_DATABASE_CONN_CTX"
)

func Conn(ctx context.Context) (*pgxpool.Conn, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, errors.Wrap(
			err,
			"pgx connection pool acquire error",
		)
	}
	return conn, nil
}

func ConnCtx(ctx context.Context) (*pgxpool.Conn, error) {
	v := ctx.Value(KEY_DATABASE_CONN_CTX)
	if v, ok := v.(*pgxpool.Conn); ok {
		return v, nil
	}
	return nil, errors.New("there is no database connection in the context")
}
