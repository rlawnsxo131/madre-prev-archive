package auth

type SocialAccountService interface {
	Create(sa *SocialAccount) (string, error)
	FindOneBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error)
}

type socialAccountService struct {
	commandRepo SocialAccountCommandRepository
	queryRepo   SocialAccountQueryRepository
}

func NewSocialAccountService(
	commandRepo SocialAccountCommandRepository,
	queryRepo SocialAccountQueryRepository,
) SocialAccountService {
	return &socialAccountService{
		commandRepo,
		queryRepo,
	}
}

func (s *socialAccountService) Create(sa *SocialAccount) (string, error) {
	return s.commandRepo.Create(sa)
}

func (s *socialAccountService) FindOneBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error) {
	return s.queryRepo.FindOneBySocialIdAndProvider(socialId, provider)
}
