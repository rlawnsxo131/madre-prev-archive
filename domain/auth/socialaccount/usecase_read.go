package socialaccount

import (
	"github.com/rlawnsxo131/madre-server-v2/database"
)

type ReadUseCase interface {
	FindOneBySocialIdAndProvider(params *SocialIDAndProviderDto) (*SocialAccount, error)
}

type readUseCase struct {
	repo ReadRepository
}

func NewReadUseCase(db database.Database) ReadUseCase {
	return &readUseCase{
		repo: NewReadRepository(db),
	}
}

func (uc *readUseCase) FindOneBySocialIdAndProvider(params *SocialIDAndProviderDto) (*SocialAccount, error) {
	return uc.repo.FindOneBySocialIdAndProvider(params)
}
