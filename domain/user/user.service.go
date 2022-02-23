package user

import (
	"github.com/jmoiron/sqlx"
)

type service struct {
	db *sqlx.DB
}

type UserService interface {
	FindOne(id string) (User, error)
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
