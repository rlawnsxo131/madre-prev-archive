package auth

import "github.com/jmoiron/sqlx"

type SocialAccountService interface {
	FindOneBySocialId(socialId string) (SocialAccount, error)
}

type socialAccountService struct {
	db *sqlx.DB
}

func NewSocialAccountService(db *sqlx.DB) SocialAccountService {
	return &socialAccountService{
		db: db,
	}
}

func (s *socialAccountService) FindOneBySocialId(socialId string) (SocialAccount, error) {
	socialAccountRepo := NewSocialAccountRepository(s.db)
	socialAccount, err := socialAccountRepo.FindOneBySocialId(socialId)
	return socialAccount, err
}
