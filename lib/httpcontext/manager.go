package httpcontext

import (
	"context"
	"errors"

	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/token"
)

const (
	Key_Database    = "Database"
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

func (m *manager) Database() (database.Database, error) {
	v := m.ctx.Value(Key_Database)
	if v, ok := v.(database.Database); ok {
		return v, nil
	}
	return nil, errors.New("Database is not exist")
}

func (m *manager) SetDatabase(db database.Database) context.Context {
	return context.WithValue(
		m.ctx,
		Key_Database,
		db,
	)
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
