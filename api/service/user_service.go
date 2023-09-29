package service

import (
	"context"

	"github.com/rlawnsxo131/madre-server/domain/persistence"
	"github.com/rlawnsxo131/madre-server/domain/persistence/repository"
)

type UserService struct {
	conn     persistence.Conn
	userRepo *repository.UserRepository
}

func NewUserService(
	conn persistence.Conn,
) *UserService {
	return &UserService{
		conn:     conn,
		userRepo: repository.NewUserRepository(),
	}
}

func (us *UserService) IsExistsUsername(username string) (bool, error) {
	exists, err := us.userRepo.ExistsByUsername(
		context.Background(),
		&persistence.QueryOptions{
			Conn: us.conn,
		},
		username,
	)

	if err != nil {
		return false, err
	}

	return exists, err
}
