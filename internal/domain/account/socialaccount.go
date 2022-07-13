package account

import (
	"time"
)

type SocialAccount struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	SocialID  string    `json:"social_id"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
