package data

import "github.com/rlawnsxo131/madre-server-v2/database"

type ReadUseCase interface {
	FindAll(limit int) ([]*Data, error)
	FindOneById(id string) (*Data, error)
}

type readUseCase struct {
	repo ReadRepository
}

func NewReadUseCase(db database.Database) ReadUseCase {
	return &readUseCase{
		repo: NewReadRepository(db),
	}
}

func (uc *readUseCase) FindAll(limit int) ([]*Data, error) {
	return uc.repo.FindAll(limit)
}

func (uc *readUseCase) FindOneById(id string) (*Data, error) {
	return uc.repo.FindOneById(id)
}
