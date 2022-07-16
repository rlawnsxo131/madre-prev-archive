package repository

import "github.com/rlawnsxo131/madre-server-v3/internal/domain/user"

type userRepository struct{}

func NewUserRepository() user.UserRepository {
	return &userRepository{}
}

func (ur *userRepository) Save(u *user.User) (string, error)
