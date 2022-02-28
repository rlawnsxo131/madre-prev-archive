package auth

import (
	"database/sql"
)

type AuthService interface {
	GetIsExistSocialAccountMap(socialAccount SocialAccount, err error) (map[string]bool, error)
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) GetIsExistSocialAccountMap(socialAccount SocialAccount, err error) (map[string]bool, error) {
	exist := false

	if err != nil {
		if err != sql.ErrNoRows {
			exist = false
		} else {
			return nil, err
		}
	}
	if socialAccount.ID != 0 {
		exist = true
	}

	return map[string]bool{"exist": exist}, nil
}
