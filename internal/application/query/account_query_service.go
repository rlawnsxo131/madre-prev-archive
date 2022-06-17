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

func (aqs *accountQueryService) GetUserById(userId string) (*account.User, error) {
	return aqs.repo.FindUserById(userId)
}

func (aqs *accountQueryService) GetUserByUsername(username string) (*account.User, error) {
	return aqs.repo.FindUserByUsername(username)
}

func (aqs *accountQueryService) ExistsUserByUsername(username string) (bool, error) {
	return aqs.repo.ExistsUserByUsername(username)
}

func (aqs *accountQueryService) GetSocialAccountBySocialIdAndProvider(socialId, provider string) (*account.SocialAccount, error) {
	return aqs.repo.FindSocialAccountBySocialIdAndProvider(socialId, provider)
}

func (aqs *accountQueryService) ExistSocialAccountBySocialIdAndProvider(socialId, provider string) (bool, error) {
	return aqs.repo.ExistsSocialAccountBySocialIdAndProvider(socialId, provider)
}
