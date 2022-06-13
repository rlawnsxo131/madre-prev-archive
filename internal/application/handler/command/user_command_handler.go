package command

import (
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/common"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/persistence/repository"
	queryrepository "github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/persistence/repository/query"
)

type UserCommandHandler interface {
	CreateUser(cmd *CreateUserCommand) (*user.User, *common.MadreError)
}

type userCommandHandler struct {
	userRepository    user.UserRepository
	userDomainService user.UserDomainService
}

func NewUserCommandHandler() UserCommandHandler {
	return &userCommandHandler{
		repository.NewUserRepository(),
		user.NewUserDomainService(
			queryrepository.NewUserQueryRepository(),
			queryrepository.NewSocialAccountQueryRepository(),
		),
	}
}

func (uch *userCommandHandler) CreateUser(cmd *CreateUserCommand) (*user.User, *common.MadreError) {
	u, err := user.NewUserWithoutId(
		cmd.Email,
		cmd.Username,
		cmd.PhotoUrl,
	)
	if err != nil {
		return nil, common.NewMadreError(err, "이름을 다시 확인해 주세요.")
	}

	// socialaccount is essential when creating a new user
	err = u.SetNewSocialAccount(
		cmd.SocialId,
		cmd.SocialUsername,
		cmd.Provider,
	)
	if err != nil {
		return nil, common.NewMadreError(err)
	}

	if err := uch.userDomainService.CheckConflictUsername(u.Username); err != nil {
		return nil, err
	}
	if err := uch.userDomainService.CheckConflictSocialAccount(
		u.SocialAccount.SocialId,
		u.SocialAccount.SocialUsername,
	); err != nil {
		return nil, err
	}

	return u, nil
}
