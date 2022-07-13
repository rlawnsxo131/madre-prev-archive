package account

import (
	"time"
)

type User struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	OriginName *string   `json:"origin_name"`
	Username   *Username `json:"username"`
	PhotoUrl   *string   `json:"photo_url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
