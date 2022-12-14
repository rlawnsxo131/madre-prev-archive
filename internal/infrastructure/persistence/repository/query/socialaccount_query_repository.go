package queryrepository

import "github.com/rlawnsxo131/madre-server-v3/internal/domain/user"

type socialaccountQueryRepository struct{}

func NewSocialAccountQueryRepository() user.SocialAccountQueryRepository {
	return &socialaccountQueryRepository{}
}

func (sqr *socialaccountQueryRepository) FindByUserId(userId string) (*user.SocialAccount, error)
func (sqr *socialaccountQueryRepository) FindBySocialIdAndProvider(socialId, provider string) (*user.SocialAccount, error)
func (sqr *socialaccountQueryRepository) ExistsBySocialIdAndProvider(socialId, provider string) (bool, error)
