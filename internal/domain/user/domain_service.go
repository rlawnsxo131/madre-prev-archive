package user

import "github.com/rlawnsxo131/madre-server-v3/internal/domain/common"

type UserDomainService interface {
	CheckConflictUsername(username string) *common.MadreError
	CheckConflictSocialAccount(socialId, provider string) *common.MadreError
}

type userDomainService struct {
	userQueryRepository          UserQueryRepository
	socialaccountQueryRepository SocialAccountQueryRepository
}

func NewUserDomainService(
	userQueryRepository UserQueryRepository,
	socialaccountQueryRepository SocialAccountQueryRepository,
) UserDomainService {
	return &userDomainService{
		userQueryRepository,
		socialaccountQueryRepository,
	}
}

func (uds *userDomainService) CheckConflictUsername(username string) *common.MadreError {
	exist, err := uds.userQueryRepository.ExistsByUsername(username)
	if err != nil {
		return common.NewMadreError(err)
	}
	if exist {
		return common.NewMadreError(
			common.ErrConflictUniqValue,
			"중복된 이름입니다.",
		)
	}
	return nil
}

func (uds *userDomainService) CheckConflictSocialAccount(socialId, provider string) *common.MadreError {
	exist, err := uds.socialaccountQueryRepository.ExistsBySocialIdAndProvider(
		socialId,
		provider,
	)
	if err != nil {
		return common.NewMadreError(err)
	}
	if exist {
		return common.NewMadreError(
			common.ErrConflictUniqValue,
			"이미 가입한 소셜 계정입니다.",
		)
	}
	return nil
}
