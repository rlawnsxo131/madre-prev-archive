package command

import "github.com/rlawnsxo131/madre-server-v3/internal/domain/account"

type accountCommandService struct {
	repo account.AccountCommandRepository
}

func NewAccountCommandService(repo account.AccountCommandRepository) account.AccountCommandService {
	return &accountCommandService{repo}
}

func (s *accountCommandService) SaveAccount(ac *account.Account) (*account.Account, error) {
	_, err := s.repo.InsertUser(ac.User())
	if err != nil {
		return nil, err
	}
	ac.SocialAccount().UserID = ac.User().ID
	_, err = s.repo.InsertSocialAccount(ac.SocialAccount())
	if err != nil {
		return nil, err
	}
	return ac, nil
}
