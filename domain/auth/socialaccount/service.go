package socialaccount

import "github.com/jmoiron/sqlx"

type Service interface {
	ReadRepository
	WriteRepository
}

type service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) Service {
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
	socialAccountReadRepo := NewReadRepository(s.db)
	socialAccount, err := socialAccountReadRepo.FindOneByProviderWithSocialId(provider, socialId)
	return socialAccount, err
}
