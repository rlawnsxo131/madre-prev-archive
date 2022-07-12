package account

import (
	"database/sql"
	"regexp"
	"time"

	"github.com/pkg/errors"
)

var (
	ErrInvalidUsername       = errors.New("username is required")
	ErrInvalidEmail          = errors.New("email is required")
	ErrInvalidUsernameFormat = errors.New("username regex MatchString error")
)

type User struct {
	ID         string         `db:"id"`
	Email      string         `db:"email"`
	OriginName sql.NullString `db:"origin_name"`
	Username   string         `db:"username"`
	PhotoUrl   sql.NullString `db:"photo_url"`
	CreatedAt  time.Time      `db:"created_at"`
	UpdatedAt  time.Time      `db:"updated_at"`
}

func NewUser(
	email string,
	originName sql.NullString,
	username string,
	photoUrl sql.NullString,
) (*User, error) {
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
