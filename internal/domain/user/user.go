package user

import (
	"regexp"
	"time"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/common"
	"github.com/rlawnsxo131/madre-server-v3/utils"
)

type User struct {
	Id            string         `json:"id"`
	Email         string         `json:"email"`
	Username      string         `json:"username"`
	PhotoUrl      string         `json:"photoUrl,omitempty"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	SocialAccount *SocialAccount `json:"socialAccount,omitempty"`
}

func NewUserWithoutId(email, username, photoUrl string) (*User, error) {
	if email == "" || username == "" {
		return nil, common.ErrMissingRequiredValue
	}
	if err := validateUsername(username); err != nil {
		return nil, err
	}
	u := User{
		Email:    email,
		Username: username,
		PhotoUrl: photoUrl,
	}
	return &u, nil
}

func (u *User) SetNewSocialAccount(socialId, socialUsername, provider string) error {
	if socialId == "" || socialUsername == "" || provider == "" {
		return common.ErrMissingRequiredValue
	}
	if err := u.isSupportSocialProvider(provider); err != nil {
		return err
	}
	sa := SocialAccount{
		SocialId:       socialId,
		SocialUsername: socialUsername,
		Provider:       provider,
	}
	u.SocialAccount = &sa
	return nil
}

func (u *User) SetSocialAccount(sa *SocialAccount) {
	u.SocialAccount = sa
}

func (u *User) isSupportSocialProvider(provider string) error {
	isContain := utils.Contains(
		[]string{SOCIAL_PROVIDER_GOOGLE},
		provider,
	)
	if !isContain {
		return errors.Wrap(
			common.ErrNotSupportValue,
			"not support provider",
		)
	}
	return nil
}

func validateUsername(username string) error {
	match, err := regexp.MatchString(
		"^[a-zA-Z0-9]{1,20}$",
		username,
	)
	if err != nil {
		return errors.Wrap(
			err,
			"username regex MatchString error",
		)
	}
	if !match {
		return errors.Wrap(
			common.ErrUnprocessableValue,
			"username regex valdiation not matched",
		)

	}
	return nil
}
