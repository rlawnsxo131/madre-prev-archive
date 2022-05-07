package user

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

func (uc *usecase) Create(u *User) (string, error) {
	repo := NewWriteRepository(uc.db)
	id, err := repo.Create(u)
	return id, err
}

func (uc *usecase) FindOneById(id string) (*User, error) {
	repo := NewReadRepository(uc.db)
	u, err := repo.FindOneById(id)
	return u, err
}
