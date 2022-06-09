package auth

type SocialAccountService interface {
	Create(sa *SocialAccount) (string, error)
	FindOneBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error)
}

type SocialAccountRepository interface {
	Create(sa *SocialAccount) (string, error)
	FindOneBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error)
}
