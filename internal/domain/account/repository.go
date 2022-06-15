package account

type AccountCommandRepository interface {
	InsertUser(u *User) (*User, error)
	InsertSocialAccount(sa *SocialAccount) (*SocialAccount, error)
}

type AccountQueryRepository interface {
	FindUserById(id string) (*User, error)
	FindUserByUsername(username string) (*User, error)
	FindSocialAccountBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error)
}
