package token

import (
	"context"

	"github.com/rlawnsxo131/madre-server-v3/typeutil"
)

const (
	KEY_USER_PROFILE_CTX = typeutil.ContextStringKey("KEY_USER_PROFILE_CTX")
)

type profile struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	PhotoUrl string `json:"photo_url"`
}

func NewProfile(userId, username, photoUrl string) *profile {
	return &profile{userId, username, photoUrl}
}

func ProfileCtx(ctx context.Context) *profile {
	v := ctx.Value(KEY_USER_PROFILE_CTX)
	if v, ok := v.(*profile); ok {
		return v
	}
	return nil
}

func SetProfileCtx(ctx context.Context, p *profile) context.Context {
	return context.WithValue(ctx, KEY_USER_PROFILE_CTX, p)
}
