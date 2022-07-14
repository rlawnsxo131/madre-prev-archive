package account

import (
	"time"

	"github.com/pkg/errors"
)

var (
	ErrInvalidSocialId      = errors.New("socialId is required")
	ErrInvalidProvider      = errors.New("social account provider is required")
	ErrNotSupportedProvider = errors.New("not supported social account provider")
)

const (
	SOCIAL_PROVIDER_GOOGLE = "GOOGLE"
)

type SocialAccount struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	SocialID  string    `json:"social_id"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewSocialAccount(socialId, provider string) (*SocialAccount, error) {
	if socialId == "" {
		return nil, ErrInvalidSocialId
	}
	if provider == "" {
		return nil, ErrInvalidProvider
	}
	if err := isSupportProvider(provider); err != nil {
		return nil, err
	}
	return &SocialAccount{
		SocialID: socialId,
		Provider: provider,
	}, nil
}

func isSupportProvider(provider string) error {
	valid := false
	for _, v := range []string{SOCIAL_PROVIDER_GOOGLE} {
		if v == provider {
			valid = true
			break
		}
	}
	if !valid {
		return ErrNotSupportedProvider
	}
	return nil
}
