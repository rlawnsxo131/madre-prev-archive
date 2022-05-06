package socialaccount

import (
	"github.com/rlawnsxo131/madre-server-v2/database"
)

type Service interface {
	ReadRepository
	WriteRepository
}

type service struct {
	db database.Database
}

func NewService(db database.Database) Service {
	return &service{
		db: db,
	}
}

func (s *service) Create(socialAccount *SocialAccount) (string, error) {
	repo := NewWriteRepository(s.db)
	id, err := repo.Create(socialAccount)
	return id, err
}

func (s *service) FindOneByProviderWithSocialId(socialId string, provider string) (*SocialAccount, error) {
	repo := NewReadRepository(s.db)
	socialAccount, err := repo.FindOneByProviderWithSocialId(provider, socialId)
	return socialAccount, err
}
