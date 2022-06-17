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

func (as *accountQueryService) FindUserById(userId string) (*account.User, error) {
	return as.repo.FindUserById(userId)
}

func (as *accountQueryService) FindUserByUsername(username string) (*account.User, error) {
	return as.repo.FindUserByUsername(username)
}

func (as *accountQueryService) FindSocialAccountBySocialIdAndProvider(socialId, provider string) (*account.SocialAccount, error) {
	return as.repo.FindSocialAccountBySocialIdAndProvider(socialId, provider)
}
