package account

type SocialAccountCommandRepository interface {
	Create(sa *SocialAccount) (string, error)
}

type SocialAccountQueryRepository interface {
	FindOneBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error)
}
