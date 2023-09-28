package user

import (
	"github.com/rlawnsxo131/madre-server/core/funk"
	"github.com/rlawnsxo131/madre-server/domain/errz"
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
		return nil, errz.NewErrMissingRequiredValue(userId)
	}
	if socialId == "" {
		return nil, errz.NewErrMissingRequiredValue(socialId)
	}
	if provider == "" {
		return nil, errz.NewErrMissingRequiredValue(provider)
	}

	if !funk.Contains[string](_socialProviders, provider) {
		return nil, errz.NewErrNotSupportValue(provider)
	}

	return &userSocialAccount{
		SocialId: socialId,
		Provider: provider,
	}, nil
}
