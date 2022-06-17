package account

import (
	"database/sql"
	"regexp"
	"time"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/common"
)

type User struct {
	ID         string         `json:"id" db:"id"`
	Email      string         `json:"email" db:"email"`
	OriginName sql.NullString `json:"origin_name" db:"origin_name"`
	Username   string         `json:"username" db:"username"`
	PhotoUrl   sql.NullString `json:"photo_url" db:"photo_url"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" db:"updated_at"`
}

func (u *User) IsExist(err error) (bool, error) {
	return common.IsExistEntity(u.ID, err)
}

func (u *User) ValidateUsername() (bool, error) {
	match, err := regexp.MatchString("^[a-zA-Z0-9]{1,20}$", u.Username)
	if err != nil {
		return false, errors.Wrap(err, "ValidateUsername regex error")
	}
	return match, nil
}
