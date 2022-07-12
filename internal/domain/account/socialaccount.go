package account

import (
	"time"

	"github.com/pkg/errors"
)

const (
	SOCIAL_ACCOUNT_PROVIDER_GOOGLE = "GOOGLE"
)

var (
	ErrInvalidUserId         = errors.New("userId is required")
	ErrInvalidSocialId       = errors.New("socialId is required")
	ErrInvalidProvider       = errors.New("provider is required")
	ErrInvalidProviderGoogle = errors.New("provider is not google")
)

type SocialAccount struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	SocialID  string    `db:"social_id"`
	Provider  string    `db:"provider"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewSocialAccount(
	userId,
	socialId,
	provider string,
) (*SocialAccount, error) {
	if userId == "" {
		return nil, ErrInvalidUserId
	}
	if socialId == "" {
		return nil, ErrInvalidSocialId
	}
	if provider == "" {
		return nil, ErrInvalidProvider
	}
	if provider != SOCIAL_ACCOUNT_PROVIDER_GOOGLE {
		return nil, ErrInvalidProviderGoogle
	}

	return &SocialAccount{
		UserID:   userId,
		SocialID: socialId,
		Provider: provider,
	}, nil
}
