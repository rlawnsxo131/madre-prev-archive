package queryservice

import (
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/account"
	queryrepository "github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/repository/query"
)

type accountQueryService struct {
	repo account.AccountQueryRepository
}

func NewAccountQueryService(db rdb.Database) account.AccountQueryService {
	return &accountQueryService{
		queryrepository.NewAccountQueryRepository(db),
	}
}

func (aqs *accountQueryService) GetAccountByUserId(userId string) (*account.Account, error) {
	u, err := aqs.repo.FindUserById(userId)
	if err != nil {
		return nil, err
	}
	sa, err := aqs.repo.FindSocialAccountByUserId(u.ID)
	if err != nil {
		return nil, err
	}

	ac := account.Account{
		User:          u,
		SocialAccount: sa,
	}

	return &ac, nil
}

func (aqs *accountQueryService) GetUserById(userId string) (*account.User, error) {
	return aqs.repo.FindUserById(userId)
}

func (aqs *accountQueryService) GetUserByUsername(username string) (*account.User, error) {
	return aqs.repo.FindUserByUsername(username)
}

func (aqs *accountQueryService) GetExistsUserByUsername(username string) (bool, error) {
	return aqs.repo.ExistsUserByUsername(username)
}

func (aqs *accountQueryService) GetSocialAccountBySocialIdAndProvider(socialId, provider string) (*account.SocialAccount, error) {
	return aqs.repo.FindSocialAccountBySocialIdAndProvider(socialId, provider)
}

func (aqs *accountQueryService) GetExistsSocialAccountBySocialIdAndProvider(socialId, provider string) (bool, error) {
	return aqs.repo.ExistsSocialAccountBySocialIdAndProvider(socialId, provider)
}
