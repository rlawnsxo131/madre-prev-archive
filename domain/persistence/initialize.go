package persistence

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	_, _b, _, _ = runtime.Caller(0)
	_basePath   = filepath.Dir(_b)
)

func ExcuteInitSQL(db *sql.DB) error {
	file, err := os.ReadFile(
		filepath.Join(_basePath, "init.sql"),
	)
	if err != nil {
		return err
	}

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	queries := strings.Split(string(file), "\n\n")
	for _, query := range queries {
		if _, err := tx.Exec(query); err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()

	return nil
}
