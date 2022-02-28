package user

import (
	"time"
)

type User struct {
	ID          uint      `json:"id" db:"id"`
	UUID        string    `json:"uuid" db:"uuid"`
	AuthId      uint      `json:"auth_id" db:"auth_id"`
	Email       string    `json:"email" db:"email"`
	Username    string    `json:"username" db:"username"`
	DisplayName string    `json:"display_name" db:"display_name"`
	PhotoUrl    string    `json:"photo_url" db:"photo_url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
