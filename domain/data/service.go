package data

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

func NewService(db database.Database) *service {
	return &service{
		db: db,
	}
}

func (s *service) FindAll(limit int) ([]*Data, error) {
	readRepo := NewReadRepository(s.db)
	dataList, err := readRepo.FindAll(limit)
	return dataList, err
}

func (s *service) FindOneById(id string) (*Data, error) {
	readRepo := NewReadRepository(s.db)
	data, err := readRepo.FindOneById(id)
	return data, err
}
