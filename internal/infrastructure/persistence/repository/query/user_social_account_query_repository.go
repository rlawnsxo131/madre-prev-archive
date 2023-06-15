package queryrepository

import (
	"github.com/rlawnsxo131/madre-server-v3/core/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/persistence/mapper"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/persistence/model"
)

type userSocialAccountQueryRepository struct {
	db     rdb.SingletonDatabase
	mapper mapper.UserSocialAccountMapper
}

func NewUserSocialAccountQueryRepository(db rdb.SingletonDatabase) user.UserSocialAccountQueryRepository {
	return &userSocialAccountQueryRepository{
		db:     db,
		mapper: mapper.UserSocialAccountMapper{},
	}
}

func (sqr *userSocialAccountQueryRepository) FindByUserId(userId string) (*user.UserSocialAccount, error) {
	sa := model.SocialAccount{}

	return sqr.mapper.MapToEntity(&sa), nil
}

func (sqr *userSocialAccountQueryRepository) ExistsBySocialIdAndProvider(socialId, provider string) (bool, error) {
	return false, nil
}
