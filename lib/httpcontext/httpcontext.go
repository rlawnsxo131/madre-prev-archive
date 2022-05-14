package httpcontext

import (
	"context"

	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/lib/token"
)

const (
	Key_Database    = "Database"
	Key_HTTPLogger  = "HTTPLogger"
	Key_UserProfile = "UserProfile"
)

func HTTPLogger(ctx context.Context) logger.HTTPLogger {
	v := ctx.Value(Key_HTTPLogger)
	if v, ok := v.(logger.HTTPLogger); ok {
		return v
	}
	return nil
}

func SetHTTPLogger(ctx context.Context, l logger.HTTPLogger) context.Context {
	return context.WithValue(
		ctx,
		Key_HTTPLogger,
		l,
	)
}

func UserProfile(ctx context.Context) *token.UserProfile {
	v := ctx.Value(Key_UserProfile)
	if v, ok := v.(*token.UserProfile); ok {
		return v
	}
	return nil
}

func SetUserProfile(ctx context.Context, p *token.UserProfile) context.Context {
	return context.WithValue(
		ctx,
		Key_UserProfile,
		p,
	)
}
