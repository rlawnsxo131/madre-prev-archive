package commandservice

import (
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/account"
	commandrepository "github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/repository/command"
)

type accountCommandService struct {
	repo account.AccountCommandRepository
}

func NewAccountCommandService(db rdb.Database) account.AccountCommandService {
	return &accountCommandService{
		commandrepository.NewAccountCommandRepository(db),
	}
}

func (acs *accountCommandService) SaveAccount(u *account.User, sa *account.SocialAccount) (*account.Account, error) {
	userId, err := acs.repo.InsertUser(u)
	if err != nil {
		return nil, err
	}
	u.ID = userId

	sa.UserID = userId
	socialAccountId, err := acs.repo.InsertSocialAccount(sa)
	if err != nil {
		return nil, err
	}
	sa.ID = socialAccountId

	ac := account.Account{
		User:          u,
		SocialAccount: sa,
	}

	return &ac, nil
}
