package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
)

type WriteRepository interface {
	Create(u User) (string, error)
}

type writeRepository struct {
	ql logger.QueryLogger
}

func NewWriteRepository(db *sqlx.DB) WriteRepository {
	return &writeRepository{
		ql: logger.NewQueryLogger(db),
	}
}

func (r *writeRepository) Create(u User) (string, error) {
	var id string
	var query = "INSERT INTO public.user(email, origin_name, display_name, photo_url) VALUES(:email, :origin_name, :display_name, :photo_url) RETURNING id"

	err := r.ql.PrepareNamedGet(&id, query, u)
	if err != nil {
		return "", errors.Wrap(err, "UserWriteRepository: create")
	}

	return id, nil
}
