package account

type Account struct {
	User          *User          `json:"user"`
	SocialAccount *SocialAccount `json:"social_account"`
}
