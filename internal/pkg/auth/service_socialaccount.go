package auth

type socialAccountService struct {
	repo SocialAccountRepository
}

func NewSocialAccountService(repo SocialAccountRepository) SocialAccountService {
	return &socialAccountService{repo}
}

func (s *socialAccountService) Create(sa *SocialAccount) (string, error) {
	return s.repo.Create(sa)
}

func (s *socialAccountService) FindOneBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error) {
	return s.repo.FindOneBySocialIdAndProvider(socialId, provider)
}
