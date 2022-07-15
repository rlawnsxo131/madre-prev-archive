package entity

import "github.com/pkg/errors"

const (
	SOCIAL_PROVIDER_GOOGLE = "GOOGLE"
)

var (
	ErrInvalidSocialId      = errors.New("socialId is required")
	ErrInvalidProvider      = errors.New("social account provider is required")
	ErrNotSupportedProvider = errors.New("not supported social account provider")
)

type SocialAccount struct {
	Id             string `json:"id"`
	UserId         string `json:"user_id"`
	SocialId       string `json:"social_id"`
	SocialUsername string `json:"social_username"`
	Provider       string `json:"provider"`
}

func NewSignUpSocialAccount(socialId, socialUsername, provider string) (*SocialAccount, error) {
	if socialId == "" {
		return nil, ErrInvalidSocialId
	}
	if provider == "" {
		return nil, ErrInvalidProvider
	}
	if err := isSupportProvider(provider); err != nil {
		return nil, err
	}

	sa := SocialAccount{
		SocialId:       socialId,
		SocialUsername: socialUsername,
		Provider:       provider,
	}
	return &sa, nil
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
