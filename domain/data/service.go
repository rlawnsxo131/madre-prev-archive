package data

import (
	"github.com/jmoiron/sqlx"
)

type DataService interface {
	ReadRepository
	WriteRepository
}

type dataService struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *dataService {
	return &dataService{
		db: db,
	}
}

func (s *dataService) FindAll(limit int) ([]*Data, error) {
	dataReadRepo := NewReadRepository(s.db)
	dataList, err := dataReadRepo.FindAll(limit)
	return dataList, err
}

func (s *dataService) FindOneById(id string) (*Data, error) {
	dataReadRepo := NewReadRepository(s.db)
	data, err := dataReadRepo.FindOneById(id)
	return data, err
}
