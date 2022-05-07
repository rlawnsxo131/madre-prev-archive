package httpcontext

import (
	"context"

	"github.com/rlawnsxo131/madre-server-v2/lib/token"
)

const (
	Key_UserProfile string = "UserProfile"
)

type contextManager struct {
	ctx context.Context
}

func NewContextManager(ctx context.Context) *contextManager {
	return &contextManager{
		ctx: ctx,
	}
}

func (hcm *contextManager) UserProfile() *token.UserProfile {
	v := hcm.ctx.Value(Key_UserProfile)
	if v, ok := v.(*token.UserProfile); ok {
		return v
	}
	return nil
}

func (hcm *contextManager) SetUserProfile(p *token.UserProfile) context.Context {
	return context.WithValue(
		hcm.ctx,
		Key_UserProfile,
		p,
	)
}
