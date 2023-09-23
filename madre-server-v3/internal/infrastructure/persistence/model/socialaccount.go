package model

import (
	"database/sql"
	"time"
)

type SocialAccount struct {
	Id             string         `db:"id"`
	UserId         string         `db:"user_id"`
	SocialId       string         `db:"social_id"`
	SocialUsername sql.NullString `db:"social_username"`
	Provider       string         `db:"provider"`
	CreatedAt      time.Time      `db:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at"`
}
