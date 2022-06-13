package rdb

import (
	"context"
	"os"

	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ExcuteInitSQL(pool *pgxpool.Pool) error {
	file, err := os.ReadFile("./core/datastore/rdb/init.sql")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	queries := strings.Split(string(file), "\n\n")
	for _, query := range queries {
		if _, err := tx.Exec(ctx, query); err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	tx.Commit(ctx)

	return nil
}
