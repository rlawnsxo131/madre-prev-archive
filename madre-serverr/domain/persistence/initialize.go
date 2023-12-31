package persistence

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/rlawnsxo131/madre-server/core/of"
)

var (
	_, _b, _, _ = runtime.Caller(0)
	_basePath   = filepath.Dir(_b)
)

func ExcuteInitSQL(conn Conn) error {
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
		of.DefaultLogger().Debug().Msgf("exec query: %s", query)
		if _, err := conn.ExecContext(ctx, query); err != nil {
			return err
		}
	}

	return nil
}
