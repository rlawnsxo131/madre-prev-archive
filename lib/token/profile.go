package token

import "context"

const (
	KEY_USER_PROFILE_CTX = "KEY_USER_PROFILE_CTX"
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
