package socialaccount

type ReadUseCase interface {
	FindOneBySocialIdAndProvider(params *SocialIDAndProviderDto) (*SocialAccount, error)
}

type WriteUseCase interface {
	Create(socialAccount *SocialAccount) (string, error)
}

type ReadRepository interface {
	FindOneBySocialIdAndProvider(params *SocialIDAndProviderDto) (*SocialAccount, error)
}

type WriteRepository interface {
	Create(socialAccount *SocialAccount) (string, error)
}
