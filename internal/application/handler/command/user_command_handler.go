package command

import (
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/common"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
)

type UserCommandHandler interface {
	CreateUser(cmd *CreateUserCommand) (*user.User, error)
}

type userCommandHandler struct {
	userRepository               user.UserRepository
	userQueryRepository          user.UserQueryRepository
	socialaccountQueryRepository user.SocialAccountQueryRepository
}

func NewUserCommandHandler() UserCommandHandler {
	return &userCommandHandler{}
}

func (uch *userCommandHandler) CreateUser(cmd *CreateUserCommand) (*user.User, error) {
	u, err := user.NewSignUpUser(
		cmd.Email,
		cmd.Username,
		cmd.PhotoUrl,
		cmd.SocialId,
		cmd.SocialUsername,
		cmd.Provider,
	)
	if err != nil {
		return nil, err
	}

	exist, err := uch.userQueryRepository.ExistsByUsername(u.Username)
	if err != nil {
		return nil, err
	} else if exist {
		return nil, common.ErrConflictUniqValue
	}

	exist, err = uch.socialaccountQueryRepository.ExistsBySocialIdAndProvider(
		u.SocialAccount.SocialId,
		u.SocialAccount.Provider,
	)
	if err != nil {
		return nil, err
	} else if exist {
		return nil, common.ErrConflictUniqValue
	}

	return u, nil
}
