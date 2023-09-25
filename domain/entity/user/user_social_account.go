package user

import (
	"github.com/rlawnsxo131/madre-server/core/funk"
	"github.com/rlawnsxo131/madre-server/domain/domainerr"
)

var (
	_socialProviders = []string{"GOOGLE"}
)

type userSocialAccount struct {
	Id             int64  `json:"id"`
	UserId         int64  `json:"userId"`
	SocialId       string `json:"socialId"`
	SocialUsername string `json:"socialUsername,omitempty"`
	Provider       string `json:"provider"`
}

func newUserSocialAccount(userId int64, socialId, provider string) (*userSocialAccount, error) {
	if userId == 0 {
		return nil, domainerr.NewErrMissingRequiredValue(userId)
	}
	if socialId == "" {
		return nil, domainerr.NewErrMissingRequiredValue(socialId)
	}
	if provider == "" {
		return nil, domainerr.NewErrMissingRequiredValue(provider)
	}

	if !funk.Contains[string](_socialProviders, provider) {
		return nil, domainerr.NewErrNotSupportValue(provider)
	}

	return &userSocialAccount{
		SocialId: socialId,
		Provider: provider,
	}, nil
}
