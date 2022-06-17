package account

type Account struct {
	user          *User
	socialAccount *SocialAccount
}

func (ac *Account) User() *User {
	return ac.user
}

func (ac *Account) AddUser(u *User) {
	ac.user = u
}

func (ac *Account) SocialAccount() *SocialAccount {
	return ac.socialAccount
}

func (ac *Account) AddSocialAccount(sa *SocialAccount) {
	ac.socialAccount = sa
}
