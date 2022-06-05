package user

import "github.com/rlawnsxo131/madre-server-v2/database"

type readUseCase struct {
	repo ReadRepository
}

func NewReadUseCase(db database.Database) ReadUseCase {
	return &readUseCase{
		repo: NewReadRepository(db),
	}
}

func (uc *readUseCase) FindOneById(id string) (*User, error) {
	return uc.repo.FindOneById(id)
}

func (uc *readUseCase) FindOneByUsername(username string) (*User, error) {
	return uc.repo.FindOneByUsername(username)
}
