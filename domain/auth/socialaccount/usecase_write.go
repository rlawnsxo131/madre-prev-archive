package socialaccount

import "github.com/rlawnsxo131/madre-server-v2/database"

type WriteUseCase interface {
	WriteRepository
}

type writeUseCase struct {
	repo WriteRepository
}

func NewWriteUseCase(db database.Database) WriteUseCase {
	return &writeUseCase{
		repo: NewWriteRepository(db),
	}
}

func (uc *writeUseCase) Create(s *SocialAccount) (string, error) {
	return uc.repo.Create(s)
}
