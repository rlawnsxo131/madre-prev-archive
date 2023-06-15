package command

import (
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/common"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
)

type UserCommandHandler interface {
	CreateForSocial(cmd *CreateUserCommand) (*user.User, *common.DomainError)
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

func (uch *userCommandHandler) CreateForSocial(cmd *CreateUserCommand) (*user.User, *common.DomainError) {
	u, err := user.NewUserWithoutId(
		cmd.Email,
		cmd.Username,
		cmd.PhotoUrl,
	)
	if err != nil {
		return nil, common.NewDomainError(err)
	}

	// socialAccount is essential when creating a new user
	err = u.SetNewSocialAccount(
		cmd.SocialId,
		cmd.SocialUsername,
		cmd.Provider,
	)
	if err != nil {
		return nil, common.NewDomainError(err)
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
