package query

import (
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/common"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
)

type UserQueryHandler interface {
	Get(q *GetUserQuery) (*user.User, *common.DomainError)
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

func (uqh *userQueryHandler) Get(q *GetUserQuery) (*user.User, *common.DomainError) {
	u, err := uqh.userQueryRepository.FindById(q.UserId)
	if err != nil {
		return nil, common.NewDomainError(err)
	}

	sa, err := uqh.userSocialQueryRepository.FindByUserId(q.UserId)
	if err != nil {
		return nil, common.NewDomainError(err)
	}
	if err := u.SetSocialAccount(sa); err != nil {
		return nil, common.NewDomainError(err)
	}

	u.SetSocialAccount(sa)

	return u, nil
}
