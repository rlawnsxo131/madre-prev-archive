package data

import (
	"github.com/jmoiron/sqlx"
)

type DataService interface {
	FindAll(limit int) ([]Data, error)
	FindOneById(id uint) (Data, error)
	FindOneByUUID(uuid string) (Data, error)
}

type dataService struct {
	db *sqlx.DB
}

func NewDataService(db *sqlx.DB) DataService {
	return &dataService{
		db: db,
	}
}

func (s *dataService) FindAll(limit int) ([]Data, error) {
	dataRepo := NewDataRepository(s.db)
	dataList, err := dataRepo.FindAll(limit)
	return dataList, err
}

func (s *dataService) FindOneById(id uint) (Data, error) {
	dataRepo := NewDataRepository(s.db)
	data, err := dataRepo.FindOneById(id)
	return data, err
}

func (s *dataService) FindOneByUUID(uuid string) (Data, error) {
	dataRepo := NewDataRepository(s.db)
	data, err := dataRepo.FindOneByUUID(uuid)
	return data, err
}
