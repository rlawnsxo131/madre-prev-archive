package repository

import (
	"github.com/rlawnsxo131/madre-server-v3/core/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/persistence/mapper"
)

type userRepository struct {
	db     rdb.SingletonDatabase
	mapper mapper.UserMapper
}

func NewUserRepository(db rdb.SingletonDatabase) user.UserRepository {
	return &userRepository{
		db:     db,
		mapper: mapper.UserMapper{},
	}
}

const (
	INSERT_USER_SQL = "INSERT INTO public.user" +
		"(email, origin_name, username, photo_url)" +
		" VALUES(:email, :username, :photo_url)" +
		" RETURNING id"

	INSERT_USER_SOCIAL_ACCOUNT_SQL = "INSERT INTO public.user_social_account" +
		"(user_id, social_id, social_username, provider)" +
		" VALUES(:user_id, :social_id, :social_username, :provider)" +
		" RETURNING id"
)

func (ur *userRepository) CreateForSocial(u *user.User) (string, error) {
	return "", nil
}
