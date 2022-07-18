package rdb

import (
	"context"
	"io/ioutil"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func ExcuteInitSQL(pool *pgxpool.Pool) {
	file, err := ioutil.ReadFile("./core/datastore/rdb/init.sql")
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	queries := strings.Split(string(file), "\n\n")
	for _, query := range queries {
		if _, err := tx.Exec(ctx, query); err != nil {
			tx.Rollback(ctx)
		}
	}
	tx.Commit(ctx)
}
