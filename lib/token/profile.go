package token

import (
	"context"
)

type key int

const (
	KEY_USER_PROFILE_CTX key = iota
)

type profile struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	PhotoUrl string `json:"photo_url"`
}

func NewProfile(userId, username, photoUrl string) *profile {
	return &profile{userId, username, photoUrl}
}

func Profile(ctx context.Context) *profile {
	v := ctx.Value(KEY_USER_PROFILE_CTX)
	if v, ok := v.(*profile); ok {
		return v
	}
	return nil
}

func SetProfile(ctx context.Context, p *profile) context.Context {
	return context.WithValue(
		ctx,
		KEY_USER_PROFILE_CTX,
		p,
	)
}
