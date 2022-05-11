package socialaccount

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

func (uc *usecase) Create(s *SocialAccount) (string, error) {
	repo := NewWriteRepository(uc.db)
	id, err := repo.Create(s)
	return id, err
}

func (uc *usecase) FindOneByProviderWithSocialId(provider, socialId string) (*SocialAccount, error) {
	repo := NewReadRepository(uc.db)
	sa, err := repo.FindOneByProviderWithSocialId(provider, socialId)
	return sa, err
}
