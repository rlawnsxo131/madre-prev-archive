package queryrepository

import "github.com/rlawnsxo131/madre-server-v3/internal/domain/user"

type userSocialAccountQueryRepository struct{}

func NewSocialAccountQueryRepository() user.UserSocialAccountQueryRepository {
	return &userSocialAccountQueryRepository{}
}

func (sqr *userSocialAccountQueryRepository) FindByUserId(userId string) (*user.UserSocialAccount, error)
func (sqr *userSocialAccountQueryRepository) ExistsBySocialIdAndProvider(socialId, provider string) (bool, error)
