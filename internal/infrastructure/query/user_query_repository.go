package query

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/utils"
)

type userQueryRepository struct {
	db rdb.Database
}

func NewUserQueryRepository(db rdb.Database) user.UserQueryRepository {
	return &userQueryRepository{db}
}

func (r *userQueryRepository) FindOneById(id string) (*user.User, error) {
	var u user.User

	query := "SELECT * FROM public.user" +
		" WHERE id = $1"

	err := r.db.QueryRowx(query, id).StructScan(&u)
	if err != nil {
		customError := errors.Wrap(err, "user ReadRepository FindOneById")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return &u, err
}

func (r *userQueryRepository) FindOneByUsername(username string) (*user.User, error) {
	var u user.User

	query := "SELECT * FROM public.user" +
		" WHERE username = $1"

	err := r.db.QueryRowx(query, username).StructScan(&u)
	if err != nil {
		customError := errors.Wrap(err, "user ReadRepository FindOneByUsername")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return &u, err
}
