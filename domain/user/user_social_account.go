package user

import (
	"github.com/rlawnsxo131/madre-server/core/utils/funk"
	"github.com/rlawnsxo131/madre-server/domain/common"
)

var (
	_socialProviders = []string{"GOOGLE"}
)

type UserSocialAccount struct {
	Id             string `json:"id"`
	UserId         string `json:"userId"`
	SocialId       string `json:"socialId"`
	SocialUsername string `json:"socialUsername,omitempty"`
	Provider       string `json:"provider"`
}

func NewUserSocialAccount(socialId, provider string) (*UserSocialAccount, error) {
	if socialId == "" || provider == "" {
		return nil, common.ErrMissingRequiredValue
	}

	if isContain := funk.Contains[string](_socialProviders, provider); !isContain {
		return nil, common.ErrNotSupportValue
	}

	return &UserSocialAccount{
		SocialId: socialId,
		Provider: provider,
	}, nil
}
