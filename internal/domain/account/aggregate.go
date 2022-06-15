package account

import (
	"regexp"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/domainhelper"
)

type Account struct {
	user          *User
	socialAccount *SocialAccount
}

func (ac *Account) User() *User {
	return ac.user
}

func (ac *Account) AddUser(u *User) {
	ac.user = u
}

func (ac *Account) SocialAccount() *SocialAccount {
	return ac.socialAccount
}

func (ac *Account) AddSocialAccount(sa *SocialAccount) {
	ac.socialAccount = sa
}

func (ac *Account) ValidateUsername() (bool, error) {
	match, err := regexp.MatchString("^[a-zA-Z0-9]{1,20}$", ac.user.Username)
	if err != nil {
		return false, errors.Wrap(err, "ValidateUsername regex error")
	}
	return match, nil
}

func (ac *Account) IsExistUser(err error) (bool, error) {
	exist, err := domainhelper.IsExistEntity(ac.user.ID, err)
	return exist, err
}

func (ac *Account) IsExistSocialAccount(err error) (bool, error) {
	exist, err := domainhelper.IsExistEntity(ac.socialAccount.ID, err)
	return exist, err
}
