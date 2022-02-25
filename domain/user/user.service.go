package user

import (
	"github.com/jmoiron/sqlx"
)

type UserService interface {
	FindOne(id string) (User, error)
}

type service struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) UserService {
	return &service{
		db: db,
	}
}

func (s *service) FindOne(id string) (User, error) {
	userRepo := NewUserRepository(s.db)
	user, err := userRepo.FindOneById(id)
	return user, err
}
