package auth

import "github.com/jmoiron/sqlx"

type AuthService interface {
	FindOne(id string, uuid string) (Auth, error)
}

type service struct {
	db *sqlx.DB
}

func NewAuthService(db *sqlx.DB) AuthService {
	return &service{
		db: db,
	}
}

func (s *service) FindOne(id string, uuid string) (Auth, error) {
	key := id
	if key == "" {
		key = uuid
	}
	authRepo := NewAuthRepository(s.db)
	auth, err := authRepo.FindOneById(key)
	return auth, err
}
