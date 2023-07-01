package rdb

import (
	"context"
)

// elements interfaces
type Row interface {
	Scan(dest ...any) error
}

// function interfaces
type QueryLayer interface {
	QueryRow(ctx context.Context, sql string, args ...any) Row
}

type TxQueryLayer interface {
	QueryRow(ctx context.Context, sql string, args ...any) Row
}
