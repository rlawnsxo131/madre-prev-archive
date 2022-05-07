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

func (s *service) Create(u *User) (string, error) {
	repo := NewWriteRepository(s.db)
	id, err := repo.Create(u)
	return id, err
}

func (s *service) FindOneById(id string) (*User, error) {
	repo := NewReadRepository(s.db)
	u, err := repo.FindOneById(id)
	return u, err
}
