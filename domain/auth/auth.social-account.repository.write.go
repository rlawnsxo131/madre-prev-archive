package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type socialAccountWriteRepository struct {
	db *sqlx.DB
}

func NewSocialAccountWriteRepository(db *sqlx.DB) *socialAccountWriteRepository {
	return &socialAccountWriteRepository{
		db: db,
	}
}

func (r *socialAccountWriteRepository) Create(params CreateSocialAccountParams) (int64, error) {
	query := "INSERT INTO social_account(uuid, user_id, provider, social_id) VALUES(?, ?, ?, ?)"
	result := r.db.MustExec(query, params)
	latInsertId, err := result.LastInsertId()

	if err != nil {
		err = errors.Wrap(err, "SocialAccountRepository: create error")
	}

	return latInsertId, err
}
