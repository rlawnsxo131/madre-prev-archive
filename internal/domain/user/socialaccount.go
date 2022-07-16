package user

const (
	SOCIAL_PROVIDER_GOOGLE = "GOOGLE"
)

type SocialAccount struct {
	Id             string `json:"id"`
	UserId         string `json:"user_id"`
	SocialId       string `json:"social_id"`
	SocialUsername string `json:"social_username,omitempty"`
	Provider       string `json:"provider"`
}

func NewSignUpSocialAccount(socialId, socialUsername, provider string) *SocialAccount {
	return &SocialAccount{
		SocialId:       socialId,
		SocialUsername: socialUsername,
		Provider:       provider,
	}
}
