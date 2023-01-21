package command

import (
	"context"

	"github.com/rlawnsxo131/madre-server-v3/internal/domain/common"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
)

type UserCommandHandler interface {
	Create(ctx context.Context, cmd *CreateUserCommand) (*user.User, *common.MadreError)
}

type userCommandHandler struct {
	userRepository    user.UserRepository
	userDomainService user.UserDomainService
}

func NewUserCommandHandler(
	userRepository user.UserRepository,
	userDomainService user.UserDomainService,
) UserCommandHandler {
	return &userCommandHandler{
		userRepository,
		userDomainService,
	}
}

func (uch *userCommandHandler) Create(ctx context.Context, cmd *CreateUserCommand) (*user.User, *common.MadreError) {
	u, err := user.NewUserWithoutId(
		cmd.Email,
		cmd.Username,
		cmd.PhotoUrl,
	)
	if err != nil {
		return nil, common.NewMadreError(err)
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

	// TODO: save

	return u, nil
}
