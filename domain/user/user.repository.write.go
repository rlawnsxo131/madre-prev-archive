package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
)

type UserWriteRepository interface {
	Create(u User) (int64, error)
}

type userWriteRepository struct {
	ql logger.QueryLogger
}

func NewUserWriteRepository(db *sqlx.DB) UserWriteRepository {
	return &userWriteRepository{
		ql: logger.NewQueryLogger(db),
	}
}

func (r *userWriteRepository) Create(u User) (int64, error) {
	query := "INSERT INTO user(uuid, email, origin_name, display_name, photo_url) VALUES(:uuid, :email, :origin_name, :display_name, :photo_url)"

	result, err := r.ql.NamedExec(query, u)
	if err != nil {
		return 0, errors.Wrap(err, "SocialAccountRepository: create")
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "SocialAccountRepository: create")
	}

	return lastInsertId, nil
}
