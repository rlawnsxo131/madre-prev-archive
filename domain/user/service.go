package user

import (
	"github.com/rlawnsxo131/madre-server-v2/database"
)

type Service interface {
	ReadRepository
	WriteRepository
}

type service struct {
	db database.Database
}

func NewService(db database.Database) Service {
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
