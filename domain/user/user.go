package user

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/rlawnsxo131/madre-server/domain/common"
)

var (
	ErrUsernameRegexNotMatched = errors.Join(
		common.ErrUnprocessableValue,
		errors.New("username regex valdiation not matched"),
	)
)

type User struct {
	Id            string             `json:"id"`
	Email         string             `json:"email"`
	Username      string             `json:"username"`
	PhotoUrl      string             `json:"photoUrl,omitempty"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
	SocialAccount *UserSocialAccount `json:"socialAccount,omitempty"`
}

func NewUserWithoutId(email, photoUrl string) (*User, error) {
	if email == "" || photoUrl == "" {
		return nil, common.ErrMissingRequiredValue
	}

	// initial name is generated as uuid
	return &User{
		Email:    email,
		Username: strings.ReplaceAll("uuid 만들어야 함", "-", ""),
		PhotoUrl: photoUrl,
	}, nil
}

func NewUserWithId(id, username, email, photoUrl string) (*User, error) {
	if id == "" || username == "" || email == "" {
		return nil, common.ErrMissingRequiredValue
	}

	if err := validateUsername(username); err != nil {
		return nil, err
	}

	return &User{
		Id:       id,
		Username: username,
		Email:    email,
		PhotoUrl: photoUrl,
	}, nil
}

func (u *User) SetNewSocialAccount(socialId, provider string) error {
	socialAccount, err := NewUserSocialAccount(provider, provider)
	if err != nil {
		return err
	}
	u.SocialAccount = socialAccount

	return nil
}

func (u *User) SetSocialAccount(sa *UserSocialAccount) error {
	if sa == nil {
		return common.ErrMissingRequiredValue
	}
	u.SocialAccount = sa

	return nil
}

func validateUsername(username string) error {
	match, err := regexp.MatchString(
		"^[a-zA-Z0-9]{1,25}$",
		username,
	)
	if err != nil {
		return errors.Join(
			err,
			errors.New("username regex MatchString parse error"),
		)
	}
	if !match {
		return ErrUsernameRegexNotMatched
	}
	return nil
}
