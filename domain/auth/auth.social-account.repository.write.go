package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
)

type SocialAccountWriteRepository interface {
	Create(socialAccount SocialAccount) (string, error)
}

type socialAccountWriteRepository struct {
	ql logger.QueryLogger
}

func NewSocialAccountWriteRepository(db *sqlx.DB) SocialAccountWriteRepository {
	return &socialAccountWriteRepository{
		ql: logger.NewQueryLogger(db),
	}
}

func (r *socialAccountWriteRepository) Create(socialAccount SocialAccount) (string, error) {
	query := "INSERT INTO social_account(user_id, provider, social_id) VALUES(:user_id, :provider, :social_id) RETURNING id"

	var id string
	err := r.ql.PrepareNamedGet(query, &id, socialAccount)
	if err != nil {
		return "", errors.Wrap(err, "SocialAccountWriteRepository: create")
	}

	return id, err
}
