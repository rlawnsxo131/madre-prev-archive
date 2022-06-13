package queryrepository

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/account"
	"github.com/rlawnsxo131/madre-server-v3/utils"
)

type userQueryRepository struct {
	db rdb.Database
}

func NewUserQueryRepository(db rdb.Database) account.UserQueryRepository {
	return &userQueryRepository{db}
}

func (r *userQueryRepository) FindOneById(id string) (*account.User, error) {
	var u account.User

	query := "SELECT * FROM public.user" +
		" WHERE id = $1"

	err := r.db.QueryRowx(query, id).StructScan(&u)
	if err != nil {
		customError := errors.Wrap(err, "user ReadRepository FindOneById")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return &u, err
}

func (r *userQueryRepository) FindOneByUsername(username string) (*account.User, error) {
	var u account.User

	query := "SELECT * FROM public.user" +
		" WHERE username = $1"

	err := r.db.QueryRowx(query, username).StructScan(&u)
	if err != nil {
		customError := errors.Wrap(err, "user ReadRepository FindOneByUsername")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return &u, err
}
