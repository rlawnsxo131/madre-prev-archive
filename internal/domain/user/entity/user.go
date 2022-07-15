package entity

import (
	"regexp"
	"time"

	"github.com/pkg/errors"
)

var (
	ErrInvalidEmail          = errors.New("email is required")
	ErrInvalidUsername       = errors.New("username is required")
	ErrInvalidUsernameFormat = errors.New("username regex MatchString error")
)

type User struct {
	Id            string         `json:"id"`
	Email         string         `json:"email"`
	Username      string         `json:"username"`
	PhotoUrl      string         `json:"photo_url,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	SocialAccount *SocialAccount `json:"social_account,omitempty"`
}

func NewSignUpUser(email, username, photoUrl, socialId, socialUsername, provider string) (*User, error) {
	if email == "" {
		return nil, ErrInvalidEmail
	}
	if username == "" {
		return nil, ErrInvalidUsername
	}
	if err := validateUsername(username); err != nil {
		return nil, err
	}

	sa, err := NewSignUpSocialAccount(socialId, socialUsername, provider)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:         email,
		Username:      username,
		PhotoUrl:      photoUrl,
		SocialAccount: sa,
	}, nil
}

func validateUsername(username string) error {
	match, err := regexp.MatchString(
		"^[a-zA-Z0-9]{1,20}$",
		username,
	)

	if err != nil {
		return errors.Wrap(err, "username regex MatchString error")
	}
	if !match {
		return ErrInvalidUsernameFormat
	}

	return nil
}
