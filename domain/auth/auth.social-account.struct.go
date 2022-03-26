package auth

type CreateSocialAccountParams struct {
	UserId      int64  `json:"user_id" db:"user_id"`
	UUID        string `json:"uuid" db:"uuid"`
	AccessToken string `json:"access_token" db:"access_token"`
	UserName    string `json:"username" db:"username"`
	Provider    string `json:"provider" db:"provider"`
	SocialId    string `json:"social_id" db:"social_id"`
}
