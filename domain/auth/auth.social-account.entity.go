package auth

import (
	"time"
)

const (
	Key_Provider_GOOGLE = "GOOGLE"
)

type SocialAccount struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Provider  string    `json:"provider" db:"provider"`
	SocialId  string    `json:"social_id" db:"social_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
