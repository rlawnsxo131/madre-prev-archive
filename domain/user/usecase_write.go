package user

import "github.com/rlawnsxo131/madre-server-v2/database"

type writeUseCase struct {
	repo WriteRepository
}

func NewWriteUseCase(db database.Database) WriteUseCase {
	return &writeUseCase{
		repo: NewWriteRepository(db),
	}
}

func (uc *writeUseCase) Create(u *User) (string, error) {
	return uc.repo.Create(u)
}
