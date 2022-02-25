package data

import (
	"github.com/jmoiron/sqlx"
)

type DataService interface {
	FindAll(limit int) ([]Data, error)
	FindOne(id string) (Data, error)
}

type service struct {
	db *sqlx.DB
}

func NewDataService(db *sqlx.DB) DataService {
	return &service{
		db: db,
	}
}

func (s *service) FindAll(limit int) ([]Data, error) {
	dataRepo := NewDataRepository(s.db)
	dataList, err := dataRepo.FindAll(limit)
	return dataList, err
}

func (s *service) FindOne(id string) (Data, error) {
	dataRepo := NewDataRepository(s.db)
	data, err := dataRepo.FindOneById(id)
	return data, err
}
