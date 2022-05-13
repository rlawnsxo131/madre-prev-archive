package user

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

func (uc *usecase) Create(u *User) (string, error) {
	return uc.writeRepo.Create(u)
}

func (uc *usecase) FindOneById(id string) (*User, error) {
	return uc.readRepo.FindOneById(id)
}

func (uc *usecase) FindOneByUsername(username string) (*User, error) {
	return uc.readRepo.FindOneByUsername(username)
}
