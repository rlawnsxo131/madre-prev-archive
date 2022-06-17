package account

type AccountCommandRepository interface {
	InsertUser(u *User) (*User, error)
	InsertSocialAccount(sa *SocialAccount) (*SocialAccount, error)
}

type AccountQueryRepository interface {
	FindUserById(id string) (*User, error)
	FindUserByUsername(username string) (*User, error)
	ExistsUserByUsername(username string) (bool, error)
	FindSocialAccountBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error)
	ExistsSocialAccountBySocialIdAndProvider(sodialId, provider string) (bool, error)
}
