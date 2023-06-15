package rdb

import (
	"context"
	"os"

	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func ExcuteInitSQL(db SingletonDatabase) error {
	file, err := os.ReadFile("./core/datastore/rdb/init.sql")
	if err != nil {
		return errors.Wrap(err, "ExcuteInitSQL ReadFile error")
	}

	ctx := context.Background()
	tx, err := db.Pool().BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return errors.Wrap(err, "ExcuteInitSQL BeginTx error")
	}
	defer tx.Conn().Close(ctx)

	queries := strings.Split(string(file), "\n\n")
	for _, query := range queries {
		if _, err := tx.Exec(ctx, query); err != nil {
			tx.Rollback(ctx)
			return errors.Wrap(err, "ExcuteInitSQL SQL Exec error")
		}
	}
	tx.Commit(ctx)

	return nil
}
