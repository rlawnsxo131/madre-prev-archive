package data

import "time"

type Data struct {
	ID          uint      `json:"id" db:"id"`
	UUID        string    `json:"uuid" db:"uuid"`
	UserId      uint      `json:"user_id" db:"user_id"`
	FileUrl     string    `json:"file_url" db:"file_url"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	IsPublic    bool      `json:"is_public" db:"is_public"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
