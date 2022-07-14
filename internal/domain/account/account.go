package account

type Account struct {
	user          *User
	socialAccount *SocialAccount
}

func NewAccount(email, originName, username, photoUrl, socialId, provider string) (*Account, error) {
	u, err := NewUser(email, originName, username, photoUrl)
	if err != nil {
		return nil, err
	}

	sa, err := NewSocialAccount(socialId, provider)
	if err != nil {
		return nil, err
	}

	return &Account{u, sa}, nil
}

func (ac *Account) SetUser(u *User) {
	ac.user = u
}

func (ac *Account) SetSocialAccount(sa *SocialAccount) {
	ac.socialAccount = sa
}
