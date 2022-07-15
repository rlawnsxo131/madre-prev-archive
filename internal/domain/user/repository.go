package user

import "github.com/rlawnsxo131/madre-server-v3/internal/domain/user/entity"

type UserRepository interface {
	Save(u *entity.User) (string, error)
}

type SocialAccountRepository interface {
	Save(sa *entity.SocialAccount) (string, error)
}

type UserQueryRepository interface {
	FindById(id string) (*entity.User, error)
	FindByUsername(username string) (*entity.User, error)
	ExistsByUsername(username string) (bool, error)
}

type SocialAccountQueryRepository interface {
	FindByUserId(userId string) (*entity.SocialAccount, error)
	FindBySocialIdAndProvider(socialId, provider string) (*entity.SocialAccount, error)
	ExistsBySocialIdAndProvider(socialId, provider string) (bool, error)
}
