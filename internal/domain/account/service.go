package account

type AccountCommandService interface {
	SaveAccount(ac *Account) (*Account, error)
}

type AccountQueryService interface {
	FindUserById(userId string) (*User, error)
	FindUserByUsername(username string) (*User, error)
	FindSocialAccountBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error)
}
