package user

import (
	"errors"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/rlawnsxo131/madre-server/domain/errz"
)

type User struct {
	Id            int64              `json:"id"`
	Email         string             `json:"email"`
	Username      string             `json:"username"`
	PhotoUrl      string             `json:"photoUrl,omitempty"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
	SocialAccount *userSocialAccount `json:"socialAccount,omitempty"`
}

func New(email, photoUrl string) (*User, error) {
	if email == "" {
		return nil, errz.NewErrMissingRequiredValue(email)
	}
	if photoUrl == "" {
		return nil, errz.NewErrMissingRequiredValue(photoUrl)
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errz.NewErrNotSupportValue(email)
	}
	if _, err := url.ParseRequestURI(photoUrl); err != nil {
		return nil, errz.NewErrNotSupportValue(photoUrl)
	}

	// initial name is generated as uuid
	return &User{
		Email:    email,
		Username: strings.ReplaceAll("uuid 만들어야 함", "-", ""),
		PhotoUrl: photoUrl,
	}, nil
}

func (u *User) SetUsername(username string) error {
	if username == "" {
		return errz.NewErrMissingRequiredValue(username)
	}

	if err := validateUsername(username); err != nil {
		return errz.NewErrNotSupportValue(username)
	}

	u.Username = username

	return nil
}

func (u *User) SetNewSocialAccount(socialId, provider string) error {
	socialAccount, err := newUserSocialAccount(u.Id, socialId, provider)

	if err != nil {
		return err
	}

	u.SocialAccount = socialAccount

	return nil
}

func (u *User) SetSocialAccount(sa *userSocialAccount) error {
	if sa == nil {
		return errz.NewErrMissingRequiredValue(sa)
	}

	u.SocialAccount = sa

	return nil
}

func validateUsername(username string) error {
	match, err := regexp.MatchString(
		"^[a-zA-Z0-9]{1,50}$",
		username,
	)

	if err != nil {
		return errors.Join(
			err,
			errz.NewErrUnprocessableValue(username),
		)
	}

	if !match {
		return errz.NewErrNotSupportValue(username)
	}

	return nil
}
