package user

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type readRepository struct {
	db     database.Database
	mapper entityMapper
}

func NewReadRepository(db database.Database) ReadRepository {
	return &readRepository{
		db:     db,
		mapper: entityMapper{},
	}
}

func (r *readRepository) FindOneById(id string) (*User, error) {
	var u User

	query := "SELECT * FROM public.user" +
		" WHERE id = $1"

	err := r.db.QueryRowx(query, id).StructScan(&u)
	if err != nil {
		customError := errors.Wrap(err, "user ReadRepository FindOneById")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return r.mapper.toEntity(&u), err
}

func (r *readRepository) FindOneByUsername(username string) (*User, error) {
	var u User

	query := "SELECT * FROM public.user" +
		" WHERE username = $1"

	err := r.db.QueryRowx(query, username).StructScan(&u)
	if err != nil {
		customError := errors.Wrap(err, "user ReadRepository FindOneByUsername")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return r.mapper.toEntity(&u), err
}
