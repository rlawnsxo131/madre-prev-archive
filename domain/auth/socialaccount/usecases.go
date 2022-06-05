package socialaccount

type ReadRepository interface {
	FindOneBySocialIdAndProvider(params *SocialIDAndProviderDto) (*SocialAccount, error)
}

type WriteRepository interface {
	Create(socialAccount *SocialAccount) (string, error)
}

type ReadUseCase interface {
	ReadRepository
}

type WriteUseCase interface {
	WriteRepository
}
