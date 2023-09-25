package persistence

import (
	"database/sql"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	_, _b, _, _ = runtime.Caller(0)
	_basePath   = filepath.Dir(_b)
)

func ExcuteInitSQL(tx *sql.Tx) error {
	file, err := os.ReadFile(
		filepath.Join(_basePath, "init.sql"),
	)
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

	return nil
}
