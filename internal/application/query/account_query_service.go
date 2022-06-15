package query

import (
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/account"
)

type accountQueryService struct {
	repo account.AccountQueryRepository
}

func NewAccountQueryService(repo account.AccountQueryRepository) account.AccountQueryService {
	return &accountQueryService{repo}
}

func (s *accountQueryService) FindUserById(userId string) (*account.User, error) {
	u, err := s.repo.FindUserById(userId)
	return u, err
}

func (s *accountQueryService) FindSocialAccountBySocialIdAndProvider(socialId, provider string) (*account.SocialAccount, error) {
	sa, err := s.repo.FindSocialAccountBySocialIdAndProvider(socialId, provider)
	return sa, err
}

func (s *accountQueryService) IsExistOfSocialAccount(socialId, provider string) (bool, error) {
	sa, err := s.repo.FindSocialAccountBySocialIdAndProvider(socialId, provider)
	return sa.IsExist(err)
}
