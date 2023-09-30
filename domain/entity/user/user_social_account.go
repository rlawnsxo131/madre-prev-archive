package user

import "github.com/rlawnsxo131/madre-server/core/of"

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
	var params = struct {
		UserId   int64  `validate:"required,number"`
		SocialId string `validate:"required,alphanum"`
		Provider string `validate:"required,oneof=GOOGLE"`
	}{
		UserId:   userId,
		SocialId: socialId,
		Provider: provider,
	}

	if err := of.Validator().Struct(params); err != nil {
		return nil, err
	}

	return &userSocialAccount{
		UserId:   userId,
		SocialId: socialId,
		Provider: provider,
	}, nil
}
