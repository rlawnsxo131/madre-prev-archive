package account

import (
	"time"

	"github.com/rlawnsxo131/madre-server-v3/utils"
)

type Account struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	PhotoUrl  string    `json:"photo_url"`
	SocialID  string    `json:"social_id"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAccount(u *User, sa *SocialAccount) *Account {
	return &Account{
		UserID:    u.ID,
		Username:  u.Username,
		Email:     u.Email,
		PhotoUrl:  utils.NormalizeNullString(u.PhotoUrl),
		SocialID:  sa.SocialID,
		Provider:  sa.Provider,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
