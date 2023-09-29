package service

import (
	"context"

	"github.com/rlawnsxo131/madre-server/domain/persistence"
	"github.com/rlawnsxo131/madre-server/domain/persistence/repository"
)

type UserService struct {
	db       persistence.Conn
	userRepo *repository.UserRepository
}

func NewUserService(
	db persistence.Conn,
) *UserService {
	return &UserService{
		db:       db,
		userRepo: repository.NewUserRepository(),
	}
}

func (us *UserService) IsExistsUsername(username string) (bool, error) {
	exists, err := us.userRepo.ExistsByUsername(
		context.Background(),
		username,
		&persistence.QueryOptions{
			Conn: us.db,
		},
	)

	if err != nil {
		return false, err
	}

	return exists, err
}
