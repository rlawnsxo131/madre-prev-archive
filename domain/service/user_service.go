package service

import (
	"context"

	"github.com/rlawnsxo131/madre-server/domain/domainerr"
	"github.com/rlawnsxo131/madre-server/domain/persistence"
	"github.com/rlawnsxo131/madre-server/domain/persistence/repository"
)

type UserService struct {
	db   *persistence.QueryLayer
	repo *repository.UserRepository
}

func NewUserService(db *persistence.QueryLayer, repo *repository.UserRepository) *UserService {
	return &UserService{
		db:   db,
		repo: repo,
	}
}

func (us *UserService) CheckConflictUsername(username string) *domainerr.DomainError {
	exists, err := us.repo.ExistsUsername(
		context.Background(),
		username,
		&persistence.QueryOptions{
			DB: *us.db,
		},
	)

	if err != nil {
		return domainerr.New(err)
	}
	if exists {
		return domainerr.New(
			domainerr.NewErrConflictUniqValue(username),
			"중복된 이름입니다",
		)
	}

	return nil
}
