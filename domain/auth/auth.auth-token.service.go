package auth

import "github.com/jmoiron/sqlx"

type AuthTokenService interface {
	AuthTokenReadRepository
}

type authTokenService struct {
	db *sqlx.DB
}

func NewAuthTokenService(db *sqlx.DB) AuthTokenService {
	return &authTokenService{
		db: db,
	}
}

func (s *authTokenService) FindOneByUserId(userId int64) (AuthToken, error) {
	authTokenReadRepo := NewAuthTokenReadRepository(s.db)
	authToken, err := authTokenReadRepo.FindOneByUserId(userId)
	return authToken, err
}
