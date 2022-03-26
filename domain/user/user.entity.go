package user

import (
	"database/sql"
	"time"
)

// Let's think about how to handle null string
type User struct {
	ID          int64          `json:"id" db:"id"`
	UUID        string         `json:"uuid" db:"uuid"`
	Email       string         `json:"email" db:"email"`
	Username    sql.NullString `json:"username" db:"username"`
	DisplayName string         `json:"display_name" db:"display_name"`
	PhotoUrl    sql.NullString `json:"photo_url" db:"photo_url"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}
