package repository

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
)

type userCommandRepository struct {
	db rdb.Database
}

func NewUserCommandRepository(db rdb.Database) user.UserCommandRepository {
	return &userCommandRepository{db}
}

func (r *userCommandRepository) Create(u *user.User) (string, error) {
	var id string

	query := "INSERT INTO public.user(email, origin_name, username, photo_url)" +
		" VALUES(:email, :origin_name, :username, :photo_url)" +
		" RETURNING id"

	err := r.db.PrepareNamedGet(
		&id,
		query,
		u,
	)
	if err != nil {
		return "", errors.Wrap(err, "user WriteRepository create")
	}

	return id, nil
}
