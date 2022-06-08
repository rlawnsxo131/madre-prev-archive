package auth

type SocialAccountUseCase interface {
	Create(sa *SocialAccount) (string, error)
	FindOneBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error)
}

type socialaccountUseCase struct {
	repo SocialAccountRepository
}

func NewSocialAccountUseCase(repo SocialAccountRepository) SocialAccountUseCase {
	return &socialaccountUseCase{
		repo: repo,
	}
}

func (uc *socialaccountUseCase) Create(sa *SocialAccount) (string, error) {
	return uc.repo.Create(sa)
}

func (uc *socialaccountUseCase) FindOneBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error) {
	return uc.repo.FindOneBySocialIdAndProvider(socialId, provider)
}
