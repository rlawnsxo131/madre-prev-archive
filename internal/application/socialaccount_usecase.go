package application

import "github.com/rlawnsxo131/madre-server-v3/internal/domain/account"

type socialAccountUseCase struct {
	commandRepository account.SocialAccountCommandRepository
	queryRepository   account.SocialAccountQueryRepository
}

func NewSocialAccountUseCase(
	commandRepository account.SocialAccountCommandRepository,
	queryRepository account.SocialAccountQueryRepository,
) account.SocialAccountUseCase {
	return &socialAccountUseCase{
		commandRepository,
		queryRepository,
	}
}

func (s *socialAccountUseCase) Create(sa *account.SocialAccount) (string, error) {
	return s.commandRepository.Create(sa)
}

func (s *socialAccountUseCase) FindOneBySocialIdAndProvider(socialId, provider string) (*account.SocialAccount, error) {
	return s.queryRepository.FindOneBySocialIdAndProvider(socialId, provider)
}
