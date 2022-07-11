package account

import (
	"database/sql"
	"regexp"
	"time"

	"github.com/pkg/errors"
)

var (
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

func (u *User) ValidateUsername() error {
	match, err := regexp.MatchString(
		"^[a-zA-Z0-9]{1,20}$",
		u.Username,
	)
	if err != nil {
		return errors.Wrap(err, "username regex MatchString error")
	}
	if !match {
		return ErrInvalidUsernameFormat
	}
	return nil
}

const (
	SOCIAL_PROVIDER_GOOGLE = "GOOGLE"
)

type SocialAccount struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	SocialID  string    `db:"social_id"`
	Provider  string    `db:"provider"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
