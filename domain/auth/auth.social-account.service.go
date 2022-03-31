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

func (s *socialAccountService) Create(socialAccount SocialAccount) (int64, error) {
	socialAccountWriteRepo := NewSocialAccountWriteRepository(s.db)
	lastInsertId, err := socialAccountWriteRepo.Create(socialAccount)
	return lastInsertId, err
}

func (s *socialAccountService) FindOneById(id int64) (SocialAccount, error) {
	socialAccountReadRepo := NewSocialAccountReadRepository(s.db)
	socialAccount, err := socialAccountReadRepo.FindOneById(id)
	return socialAccount, err
}

func (s *socialAccountService) FindOneBySocialId(socialId string) (SocialAccount, error) {
	socialAccountReadRepo := NewSocialAccountReadRepository(s.db)
	socialAccount, err := socialAccountReadRepo.FindOneBySocialId(socialId)
	return socialAccount, err
}
