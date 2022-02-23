package user

import (
	"time"
)

type User struct {
	Id          string    `json:"id" db:"id"`
	Email       string    `json:"email" db:"email"`
	Username    string    `json:"username" db:"username"`
	DisplayName string    `json:"display_name" db:"display_name"`
	PhotoUrl    string    `json:"photo_url" db:"photo_url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
