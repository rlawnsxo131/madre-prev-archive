package socialaccount

import "github.com/rlawnsxo131/madre-server-v2/database"

type WriteUseCase interface {
	Create(socialAccount *SocialAccount) (string, error)
}

type writeUseCase struct {
	repo WriteRepository
}

func NewWriteUseCase(db database.Database) WriteUseCase {
	return &writeUseCase{
		repo: NewWriteRepository(db),
	}
}

func (uc *writeUseCase) Create(sa *SocialAccount) (string, error) {
	return uc.repo.Create(sa)
}
