package rdb

import (
	"context"

	"github.com/pkg/errors"
)

type key int

const (
	KEY_DB_INSTANCE key = iota
)

func DBFromCtx(ctx context.Context) (*singletonDatabase, error) {
	v := ctx.Value(KEY_DB_INSTANCE)

	if v, ok := v.(*singletonDatabase); ok {
		return v, nil
	}

	return nil, errors.New("there is no database connection in the context")
}

func SetDBCtx(ctx context.Context, db *singletonDatabase) context.Context {
	return context.WithValue(
		ctx,
		KEY_DB_INSTANCE,
		db,
	)
}
