package account

type AccountRepository interface {
	InsertUser(u *User) (string, error)
	InsertSocialAccount(sa *SocialAccount) (string, error)
}

type UserQueryRepository interface {
	FindUserById(id string) (*User, error)
	FindUserByUsername(username string) (*User, error)
	ExistsUserByUsername(username string) (bool, error)
}

type SocialAccountQueryRepository interface {
	FindSocialAccountByUserId(userId string) (*SocialAccount, error)
	FindSocialAccountBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error)
	ExistsSocialAccountBySocialIdAndProvider(sodialId, provider string) (bool, error)
}
