package auth

type CreateSocialAccountParams struct {
	UUID        string `json:"uuid" db:"uuid"`
	AccessToken string `json:"access_token" db:"access_token"`
	UserName    string `json:"username" db:"username"`
	Provider    string `json:"provider" db:"provider"`
}
