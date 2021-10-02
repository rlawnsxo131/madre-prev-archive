package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID          int            `json:"id"`
	Email       string         `json:"email"`
	UserName    sql.NullString `json:"user_name"`
	DisplayName string         `json:"display_name"`
	PhotoUrl    sql.NullString `json:"photo_url"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
