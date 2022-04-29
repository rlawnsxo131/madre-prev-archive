package data

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

func NewService(db *sqlx.DB) *service {
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
