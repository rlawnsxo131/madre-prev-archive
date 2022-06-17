package account

import (
	"time"

	"github.com/rlawnsxo131/madre-server-v3/internal/domain/common"
)

const (
	SOCIAL_ACCOUNT_PROVIDER_GOOGLE = "GOOGLE"
)

type SocialAccount struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	SocialID  string    `json:"social_id" db:"social_id"`
	Provider  string    `json:"provider" db:"provider"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (sa *SocialAccount) IsExist(err error) (bool, error) {
	return common.IsExistEntity(sa.ID, err)
}
