package data

import (
	"github.com/rlawnsxo131/madre-server-v2/database"
)

type UseCase interface {
	ReadRepository
	WriteRepository
}

type usecase struct {
	readRepo  ReadRepository
	writeRepo WriteRepository
}

func NewUseCase(db database.Database) UseCase {
	return &usecase{
		readRepo:  NewReadRepository(db),
		writeRepo: NewWriteRepository(db),
	}
}

func (uc *usecase) FindAll(limit int) ([]*Data, error) {
	return uc.readRepo.FindAll(limit)
}

func (uc *usecase) FindOneById(id string) (*Data, error) {
	return uc.readRepo.FindOneById(id)
}
