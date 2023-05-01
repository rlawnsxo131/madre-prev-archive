package httplogger

import (
	"context"
	"errors"
)

type key int

const (
	KEY_LOG_ENTRY_CTX key = iota
)

func LogEntryFromCtx(ctx context.Context) (HTTPLogEntry, error) {
	v := ctx.Value(KEY_LOG_ENTRY_CTX)
	if v, ok := v.(HTTPLogEntry); ok {
		return v, nil
	}
	return nil, errors.New("there is no httpLogEntry in the context")
}

func SetLogEntryCtx(ctx context.Context, le HTTPLogEntry) context.Context {
	return context.WithValue(
		ctx,
		KEY_LOG_ENTRY_CTX,
		le,
	)
}
