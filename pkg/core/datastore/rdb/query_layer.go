package rdb

import (
	"context"
)

// elements interfaces
type Row interface {
	Scan(dest ...any) error
}

// function interfaces
type CommonQueryLayer interface {
	QueryRow(ctx context.Context, sql string, args ...any) Row
}

type QueryLayer interface {
	CommonQueryLayer
}

type TxQueryLayer interface {
	CommonQueryLayer
}
