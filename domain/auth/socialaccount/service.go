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

func (s *service) Create(socialAccount SocialAccount) (string, error) {
	writeRepo := NewWriteRepository(s.db)
	id, err := writeRepo.Create(socialAccount)
	return id, err
}

func (s *service) FindOneByProviderWithSocialId(socialId string, provider string) (*SocialAccount, error) {
	readRepo := NewReadRepository(s.db)
	socialAccount, err := readRepo.FindOneByProviderWithSocialId(provider, socialId)
	return socialAccount, err
}
