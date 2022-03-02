package auth

import "time"

type SocialAccount struct {
	ID        uint      `json:"id" db:"id"`
	UUID      string    `json:"uuid" db:"uuid"`
	UserId    uint      `json:"user_id" db:"user_id"`
	Provider  string    `json:"provider" db:"provider"`
	SocialId  string    `json:"social_id" db:"social_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
