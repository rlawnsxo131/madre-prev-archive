package auth

type SocialAccountService interface {
	Create(sa *SocialAccount) (string, error)
	FindOneBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error)
}

type socialAccountService struct {
	command SocialAccountCommandRepository
	query   SocialAccountQueryRepository
}

func NewSocialAccountService(
	command SocialAccountCommandRepository,
	query SocialAccountQueryRepository,
) SocialAccountService {
	return &socialAccountService{
		command,
		query,
	}
}

func (s *socialAccountService) Create(sa *SocialAccount) (string, error) {
	return s.command.Create(sa)
}

func (s *socialAccountService) FindOneBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error) {
	return s.query.FindOneBySocialIdAndProvider(socialId, provider)
}
