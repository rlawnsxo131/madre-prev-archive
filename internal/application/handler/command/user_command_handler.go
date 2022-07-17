package command

import (
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/common"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
)

type UserCommandHandler interface {
	CreateUser(cmd *CreateUserCommand) (*user.User, *common.MadreError)
}

type userCommandHandler struct {
	userRepository               user.UserRepository
	userQueryRepository          user.UserQueryRepository
	socialaccountQueryRepository user.SocialAccountQueryRepository
}

func NewUserCommandHandler() UserCommandHandler {
	return &userCommandHandler{}
}

func (uch *userCommandHandler) CreateUser(cmd *CreateUserCommand) (*user.User, *common.MadreError) {
	u, err := user.NewSignUpUser(
		cmd.Email,
		cmd.Username,
		cmd.PhotoUrl,
		cmd.SocialId,
		cmd.SocialUsername,
		cmd.Provider,
	)
	if err != nil {
		return nil, common.NewMadreError(err)
	}

	exist, err := uch.userQueryRepository.ExistsByUsername(u.Username)
	if err != nil {
		return nil, common.NewMadreError(err)
	} else if exist {
		return nil, common.NewMadreError(
			common.ErrConflictUniqValue,
			"중복된 이름입니다.",
		)
	}

	exist, err = uch.socialaccountQueryRepository.ExistsBySocialIdAndProvider(
		u.SocialAccount.SocialId,
		u.SocialAccount.Provider,
	)
	if err != nil {
		return nil, common.NewMadreError(err)
	} else if exist {
		return nil, common.NewMadreError(
			common.ErrConflictUniqValue,
			"이미 가입한 소셜 계정입니다.",
		)
	}

	return u, nil
}
