package command

import "github.com/rlawnsxo131/madre-server-v3/internal/domain/account"

type accountCommandService struct {
	repo account.AccountCommandRepository
}

func NewAccountCommandService(repo account.AccountCommandRepository) account.AccountCommandService {
	return &accountCommandService{repo}
}

func (acs *accountCommandService) SaveAccount(u *account.User, sa *account.SocialAccount) (*account.Account, error) {
	u, err := acs.repo.InsertUser(u)
	if err != nil {
		return nil, err
	}

	sa.UserID = u.ID
	sa, err = acs.repo.InsertSocialAccount(sa)
	if err != nil {
		return nil, err
	}

	ac := &account.Account{}
	ac.AddUser(u)
	ac.AddSocialAccount(sa)

	return ac, nil
}
