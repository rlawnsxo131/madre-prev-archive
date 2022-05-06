package user

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
)

type WriteRepository interface {
	Create(u *User) (string, error)
}

type writeRepository struct {
	db     database.Database
	mapper entityMapper
}

func NewWriteRepository(db database.Database) WriteRepository {
	return &writeRepository{
		db:     db,
		mapper: entityMapper{},
	}
}

func (r *writeRepository) Create(u *User) (string, error) {
	var id string

	query := "INSERT INTO public.user(email, origin_name, display_name, photo_url)" +
		" VALUES(:email, :origin_name, :display_name, :photo_url) RETURNING id"

	err := r.db.PrepareNamedGet(
		&id,
		query,
		r.mapper.toModel(u),
	)
	if err != nil {
		return "", errors.Wrap(err, "UserWriteRepository: create")
	}

	return id, nil
}
