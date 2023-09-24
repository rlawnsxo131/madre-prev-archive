package user

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"github.com/rlawnsxo131/madre-server/domain/domainerr"
)

type User struct {
	Id            string             `json:"id"`
	Email         string             `json:"email"`
	Username      string             `json:"username"`
	PhotoUrl      string             `json:"photoUrl,omitempty"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
	SocialAccount *userSocialAccount `json:"socialAccount,omitempty"`
}

func NewUserWithoutId(email, photoUrl string) (*User, error) {
	if email == "" {
		return nil, domainerr.NewErrMissingRequiredValue(email)
	}
	if photoUrl == "" {
		return nil, domainerr.NewErrMissingRequiredValue(photoUrl)
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return nil, domainerr.NewErrNotSupportValue(email)
	}

	// initial name is generated as uuid
	return &User{
		Email:    email,
		Username: strings.ReplaceAll("uuid 만들어야 함", "-", ""),
		PhotoUrl: photoUrl,
	}, nil
}

func NewUserWithId(id, username, email, photoUrl string) (*User, error) {
	if id == "" {
		return nil, domainerr.NewErrMissingRequiredValue(id)
	}
	if username == "" {
		return nil, domainerr.NewErrMissingRequiredValue(username)
	}
	if email == "" {
		return nil, domainerr.NewErrMissingRequiredValue(email)
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return nil, domainerr.NewErrNotSupportValue(email)
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
	socialAccount, err := newUserSocialAccount(socialId, provider)

	if err != nil {
		return err
	}

	u.SocialAccount = socialAccount

	return nil
}

func (u *User) SetSocialAccount(sa *userSocialAccount) error {
	if sa == nil {
		return domainerr.NewErrMissingRequiredValue(sa)
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
			domainerr.NewErrUnprocessableValue(username),
		)
	}

	if !match {
		return domainerr.NewErrNotSupportValue(username)
	}

	return nil
}
