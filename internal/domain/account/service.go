package account

type AccountCommandService interface {
	SaveAccount(u *User, sa *SocialAccount) (*Account, error)
}

type AccountQueryService interface {
	GetAccountByUserId(userId string) (*Account, error)
	GetUserById(userId string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	ExistsUserByUsername(username string) (bool, error)
	GetSocialAccountBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error)
	ExistsSocialAccountBySocialIdAndProvider(socialId, provider string) (bool, error)
}
