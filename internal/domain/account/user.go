package account

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
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	OriginName string    `json:"origin_name"`
	Username   string    `json:"username"`
	PhotoUrl   string    `json:"photo_url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewUser(email, originName, username, photoUrl string) (*User, error) {
	if email == "" {
		return nil, ErrInvalidEmail
	}
	if username == "" {
		return nil, ErrInvalidUsername
	}
	if err := validateUsername(username); err != nil {
		return nil, err
	}

	return &User{
		Email:      email,
		OriginName: originName,
		Username:   username,
		PhotoUrl:   photoUrl,
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
