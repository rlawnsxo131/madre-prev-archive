package repository

import (
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/persistence/mapper"
	"github.com/rlawnsxo131/madre-server-v3/pkg/core/datastore/rdb"
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

func (ur *userRepository) CreateForSocial(u *user.User) (string, error) {
	return "", nil
}
