package auth

import "time"

type AuthToken struct {
	ID           int64     `json:"id" db:"id"`
	UUID         string    `json:"uuid" db:"uuid"`
	UserId       int64     `json:"user_id" db:"user_id"`
	RefreshToken string    `json:"refresh_token" db:"refresh_token"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
