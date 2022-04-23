package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
)

type UserWriteRepository interface {
	Create(u User) (string, error)
}

type userWriteRepository struct {
	ql logger.QueryLogger
}

func NewUserWriteRepository(db *sqlx.DB) UserWriteRepository {
	return &userWriteRepository{
		ql: logger.NewQueryLogger(db),
	}
}

func (r *userWriteRepository) Create(u User) (string, error) {
	var id string
	var query = "INSERT INTO public.user(email, origin_name, display_name, photo_url) VALUES(:email, :origin_name, :display_name, :photo_url) RETURNING id"

	err := r.ql.PrepareNamedGet(&id, query, u)
	if err != nil {
		return "", errors.Wrap(err, "UserWriteRepository: create")
	}

	return id, nil
}
