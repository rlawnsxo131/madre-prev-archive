package httpcontext

import (
	"context"

	"github.com/rlawnsxo131/madre-server-v2/lib/token"
)

const (
	key_RequestID   = "RequestID"
	Key_UserProfile = "UserProfile"
)

type manager struct {
	ctx context.Context
}

func NewManager(ctx context.Context) *manager {
	return &manager{
		ctx: ctx,
	}
}

func (m *manager) RequestId() string {
	v := m.ctx.Value(key_RequestID)
	if v, ok := v.(string); ok {
		return v
	}
	return ""
}

func (m *manager) SetRequestId(id string) context.Context {
	return context.WithValue(
		m.ctx,
		key_RequestID,
		id,
	)
}

func (m *manager) UserProfile() *token.UserProfile {
	v := m.ctx.Value(Key_UserProfile)
	if v, ok := v.(*token.UserProfile); ok {
		return v
	}
	return nil
}

func (m *manager) SetUserProfile(p *token.UserProfile) context.Context {
	return context.WithValue(
		m.ctx,
		Key_UserProfile,
		p,
	)
}
