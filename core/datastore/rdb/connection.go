package rdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type key int

const (
	KEY_DATABASE_CONN_CTX key = iota
)

func Conn() (*pgxpool.Conn, error) {
	conn, err := pool.Acquire(context.Background())

	if err != nil {
		return nil, errors.Wrap(
			err,
			"pgx connection pool acquire error",
		)
	}

	return conn, nil
}

func ConnFromCtx(ctx context.Context) (*pgxpool.Conn, error) {
	v := ctx.Value(KEY_DATABASE_CONN_CTX)

	if v, ok := v.(*pgxpool.Conn); ok {
		return v, nil
	}

	return nil, errors.New("there is no database connection in the context")
}

func SetConnCtx(ctx context.Context, conn *pgxpool.Conn) context.Context {
	return context.WithValue(
		ctx,
		KEY_DATABASE_CONN_CTX,
		conn,
	)
}
