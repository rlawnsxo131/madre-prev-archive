package auth

import (
	"database/sql"
	"regexp"

	socialaccount "github.com/rlawnsxo131/madre-server-v2/domain/auth/social_account"
)

type AuthService interface {
	GetExistSocialAccountMap(socialAccount *socialaccount.SocialAccount, err error) (map[string]bool, error)
	ValidateDisplayName(userName string) (bool, error)
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) GetExistSocialAccountMap(socialAccount *socialaccount.SocialAccount, err error) (map[string]bool, error) {
	exist := false

	if err != nil {
		if err == sql.ErrNoRows {
			exist = false
		} else {
			return nil, err
		}
	}

	if socialAccount.ID != "" {
		exist = true
	}

	return map[string]bool{"exist": exist}, nil
}

func (s *authService) ValidateDisplayName(displayName string) (bool, error) {
	match, err := regexp.MatchString("^[a-zA-Z0-9가-힣]{1,16}$", displayName)
	if err != nil {
		return false, err
	}
	return match, nil
}
