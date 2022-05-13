package httpcontext

import (
	"context"

	"github.com/rlawnsxo131/madre-server-v2/lib/token"
)

const (
	Key_Database    = "Database"
	key_RequestID   = "RequestID"
	Key_UserProfile = "UserProfile"
)

func RequestId(ctx context.Context) string {
	v := ctx.Value(key_RequestID)
	if v, ok := v.(string); ok {
		return v
	}
	return ""
}

func SetRequestId(ctx context.Context, id string) context.Context {
	return context.WithValue(
		ctx,
		key_RequestID,
		id,
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
