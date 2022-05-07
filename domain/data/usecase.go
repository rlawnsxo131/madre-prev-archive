package data

import (
	"github.com/rlawnsxo131/madre-server-v2/database"
)

type UseCase interface {
	ReadRepository
	WriteRepository
}

type usecase struct {
	db database.Database
}

func NewUseCase(db database.Database) UseCase {
	return &usecase{
		db: db,
	}
}

func (uc *usecase) FindAll(limit int) ([]*Data, error) {
	repo := NewReadRepository(uc.db)
	dd, err := repo.FindAll(limit)
	return dd, err
}

func (uc *usecase) FindOneById(id string) (*Data, error) {
	repo := NewReadRepository(uc.db)
	data, err := repo.FindOneById(id)
	return data, err
}
