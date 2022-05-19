package socialaccount

import (
	"database/sql"
	"time"
)

const (
	Key_Provider_GOOGLE = "GOOGLE"
)

type SocialAccount struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	SocialId  string    `json:"social_id" db:"social_id"`
	Provider  string    `json:"provider" db:"provider"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (sa *SocialAccount) IsExist(err error) (bool, error) {
	exist := false

	if err != nil {
		if err == sql.ErrNoRows {
			return exist, nil
		} else {
			return exist, err
		}
	}

	if sa.ID != "" {
		exist = true
	}

	return exist, nil
}
