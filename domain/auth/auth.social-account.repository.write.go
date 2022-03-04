package auth

import "github.com/jmoiron/sqlx"

type SocialAccountWriteRepository interface{}

type socialAccountWriteRepository struct {
	db *sqlx.DB
}

func NewSocialAccountWriteRepository(db *sqlx.DB) SocialAccountWriteRepository {
	return &socialAccountWriteRepository{
		db: db,
	}
}
