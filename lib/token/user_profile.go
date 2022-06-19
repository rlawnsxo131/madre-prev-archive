package token

import "context"

const (
	KEY_USER_PROFILE_CTX = "KEY_USER_PROFILE_CTX"
)

type userProfile struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	PhotoUrl string `json:"photo_url"`
}

func NewUserProfile(userId, username, photoUrl string) *userProfile {
	return &userProfile{userId, username, photoUrl}
}

func UserProfileCtx(ctx context.Context) *userProfile {
	v := ctx.Value(KEY_USER_PROFILE_CTX)
	if v, ok := v.(*userProfile); ok {
		return v
	}
	return nil
}

func SetUserProfileCtx(ctx context.Context, p *userProfile) context.Context {
	return context.WithValue(ctx, KEY_USER_PROFILE_CTX, p)
}
