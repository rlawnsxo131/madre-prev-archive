package user

import (
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/common"
)

type UserDomainService interface {
	CheckConflictUsername(username string) *common.DomainError
	CheckConflictSocialAccount(socialId, provider string) *common.DomainError
}

type userDomainService struct {
	userQueryRepository          UserQueryRepository
	socialAccountQueryRepository UserSocialAccountQueryRepository
}

func NewUserDomainService(
	userQueryRepository UserQueryRepository,
	userSocialAccountQueryRepository UserSocialAccountQueryRepository,
) UserDomainService {
	return &userDomainService{
		userQueryRepository,
		userSocialAccountQueryRepository,
	}
}

func (uds *userDomainService) CheckConflictUsername(username string) *common.DomainError {
	exist, err := uds.userQueryRepository.ExistsByUsername(username)
	if err != nil {
		return common.NewDomainError(err)
	}
	if exist {
		return common.NewDomainError(
			common.ErrConflictUniqValue,
			"중복된 이름입니다.",
		)
	}
	return nil
}

func (uds *userDomainService) CheckConflictSocialAccount(socialId, provider string) *common.DomainError {
	exist, err := uds.socialAccountQueryRepository.ExistsBySocialIdAndProvider(
		socialId,
		provider,
	)
	if err != nil {
		return common.NewDomainError(err)
	}
	if exist {
		return common.NewDomainError(
			common.ErrConflictUniqValue,
			"이미 가입한 소셜 계정입니다.",
		)
	}
	return nil
}
