package socialaccount

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
)

type WriteRepository interface {
	Create(socialAccount SocialAccount) (string, error)
}

type writeRepository struct {
	ql logger.QueryLogger
}

func NewWriteRepository(db *sqlx.DB) WriteRepository {
	return &writeRepository{
		ql: logger.NewQueryLogger(db),
	}
}

func (r *writeRepository) Create(socialAccount SocialAccount) (string, error) {
	var id string
	var query = "INSERT INTO social_account(user_id, provider, social_id) VALUES(:user_id, :provider, :social_id) RETURNING id"

	err := r.ql.PrepareNamedGet(&id, query, socialAccount)
	if err != nil {
		return "", errors.Wrap(err, "writeRepository: create")
	}

	return id, err
}
