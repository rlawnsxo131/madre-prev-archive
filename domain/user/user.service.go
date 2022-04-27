package user

import (
	"github.com/jmoiron/sqlx"
)

type UserService interface {
	UserReadRepository
	UserWriteRepository
}

type userService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) UserService {
	return &userService{
		db: db,
	}
}

func (s *userService) Create(u User) (string, error) {
	userWriteRepo := NewUserWriteRepository(s.db)
	id, err := userWriteRepo.Create(u)
	return id, err
}

func (s *userService) FindOneById(id string) (*User, error) {
	userReadRepo := NewUserReadRepository(s.db)
	user, err := userReadRepo.FindOneById(id)
	return user, err
}
