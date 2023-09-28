package persistence

import (
	"context"
	"log"
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

func ExcuteInitSQL(db QueryLayer) error {
	file, err := os.ReadFile(
		filepath.Join(_basePath, "init.sql"),
	)
	if err != nil {
		return err
	}

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	queries := strings.Split(string(file), "\n\n")
	for _, query := range queries {
		log.Printf("exec query: %s", query)
		if _, err := db.ExecContext(ctx, query); err != nil {
			return err
		}
	}

	return nil
}
