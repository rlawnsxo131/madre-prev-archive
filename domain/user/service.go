package user

import (
	"github.com/jmoiron/sqlx"
)

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

func (s *service) Create(u User) (string, error) {
	userWriteRepo := NewWriteRepository(s.db)
	id, err := userWriteRepo.Create(u)
	return id, err
}

func (s *service) FindOneById(id string) (*User, error) {
	userReadRepo := NewReadRepository(s.db)
	user, err := userReadRepo.FindOneById(id)
	return user, err
}
