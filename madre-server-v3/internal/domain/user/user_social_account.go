package user

const (
	SOCIAL_PROVIDER_GOOGLE = "GOOGLE"
)

type UserSocialAccount struct {
	Id             string `json:"id"`
	UserId         string `json:"userId"`
	SocialId       string `json:"socialId"`
	SocialUsername string `json:"socialUsername,omitempty"`
	Provider       string `json:"provider"`
}
