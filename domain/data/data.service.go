package data

import (
	"github.com/jmoiron/sqlx"
)

type DataService interface {
	DataReadRepository
	DataWriteRepository
}

type dataService struct {
	db *sqlx.DB
}

func NewDataService(db *sqlx.DB) *dataService {
	return &dataService{
		db: db,
	}
}

func (s *dataService) FindAll(limit int) ([]Data, error) {
	dataReadRepo := NewDataReadRepository(s.db)
	dataList, err := dataReadRepo.FindAll(limit)
	return dataList, err
}

func (s *dataService) FindOneById(id string) (Data, error) {
	dataReadRepo := NewDataReadRepository(s.db)
	data, err := dataReadRepo.FindOneById(id)
	return data, err
}
