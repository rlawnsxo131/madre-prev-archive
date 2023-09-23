package token

import (
	"context"
)

type key int

const (
	KEY_USER_PROFILE_CTX key = iota
)

type profile struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	PhotoUrl string `json:"photo_url"`
}

func NewProfile(userId, username, photoUrl string) *profile {
	return &profile{
		userId,
		username,
		photoUrl,
	}
}

func ProfileFromCtx(ctx context.Context) *profile {
	v := ctx.Value(KEY_USER_PROFILE_CTX)
	if p, ok := v.(*profile); ok {
		return p
	}
	return nil
}

func SetProfileCtx(ctx context.Context, p *profile) context.Context {
	return context.WithValue(
		ctx,
		KEY_USER_PROFILE_CTX,
		p,
	)
}
