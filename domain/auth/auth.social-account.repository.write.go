package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
)

type SocialAccountWriteRepository interface {
	Create(socialAccount SocialAccount) (int64, error)
}

type socialAccountWriteRepository struct {
	ql logger.QueryLogger
}

func NewSocialAccountWriteRepository(db *sqlx.DB) SocialAccountWriteRepository {
	return &socialAccountWriteRepository{
		ql: logger.NewQueryLogger(db),
	}
}

func (r *socialAccountWriteRepository) Create(socialAccount SocialAccount) (int64, error) {
	query := "INSERT INTO social_account(uuid, user_id, provider, social_id) VALUES(:uuid, :user_id, :provider, :social_id)"

	result, err := r.ql.NamedExec(query, socialAccount)
	if err != nil {
		return 0, errors.Wrap(err, "SocialAccountRepository: create")
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "SocialAccountRepository: create")
	}

	return lastInsertId, nil
}
