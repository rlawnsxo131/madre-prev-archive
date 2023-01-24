package query

import (
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/common"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
)

type UserQueryHandler interface {
	Get(userId string) (*user.User, *common.MadreError)
}

type userQueryHandler struct {
	userQueryRepository       user.UserQueryRepository
	userSocialQueryRepository user.UserSocialAccountQueryRepository
}

func NewUserQueryHandler(
	userQueryRepository user.UserQueryRepository,
	userSocialAccountQueryRepository user.UserSocialAccountQueryRepository,
) UserQueryHandler {
	return &userQueryHandler{
		userQueryRepository,
		userSocialAccountQueryRepository,
	}
}

func (uqr *userQueryHandler) Get(userId string) (*user.User, *common.MadreError) {
	u, err := uqr.userQueryRepository.FindById(userId)
	if err != nil {
		return nil, common.NewMadreError(err)
	}

	sa, err := uqr.userSocialQueryRepository.FindByUserId(userId)
	if err != nil {
		return nil, common.NewMadreError(err)
	}
	if err := u.SetSocialAccount(sa); err != nil {
		return nil, common.NewMadreError(err)
	}

	return u, nil
}
