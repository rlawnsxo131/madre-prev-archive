package token

import "context"

func UserProfileCtx(ctx context.Context) *UserProfile {
	v := ctx.Value(Key_UserProfileCtx)
	if v, ok := v.(*UserProfile); ok {
		return v
	}
	return nil
}

func SetUserProfileCtx(ctx context.Context, p *UserProfile) context.Context {
	return context.WithValue(ctx, Key_UserProfileCtx, p)
}
