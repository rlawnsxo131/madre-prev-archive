package service

import (
	"context"

	"github.com/rlawnsxo131/madre-server/domain/persistence"
	"github.com/rlawnsxo131/madre-server/domain/persistence/repository"
)

type UserAccountService struct {
	db       persistence.QueryLayer
	userRepo *repository.UserRepository
}

func NewUserAccountService(
	db persistence.QueryLayer,
) *UserAccountService {
	return &UserAccountService{
		db:       db,
		userRepo: repository.NewUserRepository(),
	}
}

func (uas *UserAccountService) IsConflictUsername(username string) (bool, error) {
	exists, err := uas.userRepo.ExistsUsername(
		context.Background(),
		username,
		&persistence.QueryOptions{
			DB: uas.db,
		},
	)

	if err != nil {
		return false, err
	}

	return exists, err
}
