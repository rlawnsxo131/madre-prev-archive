package persistence

import (
	"context"
	"database/sql"
)

type QueryLayer interface {
	QueryRow(query string, args ...any) *sql.Row
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

func a() {
	// db, _ := sql.Open("", "")
	// db.QueryRowContext()
}
