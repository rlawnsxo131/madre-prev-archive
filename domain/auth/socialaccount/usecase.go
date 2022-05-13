package socialaccount

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

func (uc *usecase) Create(s *SocialAccount) (string, error) {
	return uc.writeRepo.Create(s)
}

func (uc *usecase) FindOneByProviderWithSocialId(provider, socialId string) (*SocialAccount, error) {
	return uc.readRepo.FindOneByProviderWithSocialId(provider, socialId)
}
