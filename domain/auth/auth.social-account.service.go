package auth

import "github.com/jmoiron/sqlx"

type SocialAccountService interface {
	SocialAccountReadRepository
	SocialAccountWriteRepository
}

type socialAccountService struct {
	db *sqlx.DB
}

func NewSocialAccountService(db *sqlx.DB) SocialAccountService {
	return &socialAccountService{
		db: db,
	}
}

func (s *socialAccountService) Create(socialAccount *SocialAccount) (string, error) {
	socialAccountWriteRepo := NewSocialAccountWriteRepository(s.db)
	id, err := socialAccountWriteRepo.Create(socialAccount)
	return id, err
}

func (s *socialAccountService) FindOneByProviderWithSocialId(socialId string, provider string) (*SocialAccount, error) {
	socialAccountReadRepo := NewSocialAccountReadRepository(s.db)
	socialAccount, err := socialAccountReadRepo.FindOneByProviderWithSocialId(provider, socialId)
	return socialAccount, err
}
