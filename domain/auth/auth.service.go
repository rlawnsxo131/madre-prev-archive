package auth

import (
	"database/sql"
	"regexp"
)

type AuthService interface {
	GetExistSocialAccountMap(socialAccount SocialAccount, err error) (map[string]bool, error)
	ValidateUserName(userName string) (bool, error)
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) GetExistSocialAccountMap(socialAccount SocialAccount, err error) (map[string]bool, error) {
	exist := false

	if err != nil {
		if err == sql.ErrNoRows {
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

func (s *authService) ValidateUserName(userName string) (bool, error) {
	match, err := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", userName)
	if err != nil {
		return false, err
	}
	return match, nil
}
