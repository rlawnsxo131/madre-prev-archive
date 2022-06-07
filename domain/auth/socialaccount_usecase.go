package auth

type SocialAccountUseCase interface {
	Create(sa *SocialAccount) (string, error)
	FindOneBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error)
}

type socialAccountUseCase struct {
	repo SocialAccountRepository
}

func NewSocialAccountUseCase(repo SocialAccountRepository) SocialAccountUseCase {
	return &socialAccountUseCase{
		repo: repo,
	}
}

func (uc *socialAccountUseCase) Create(sa *SocialAccount) (string, error) {
	return uc.repo.Create(sa)
}

func (uc *socialAccountUseCase) FindOneBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error) {
	return uc.repo.FindOneBySocialIdAndProvider(socialId, provider)
}
