package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type ReadRepository interface {
	FindOneById(id string) (*User, error)
}

type readRepository struct {
	ql logger.QueryLogger
}

func NewReadRepository(db *sqlx.DB) ReadRepository {
	return &readRepository{
		ql: logger.NewQueryLogger(db),
	}
}

func (r *readRepository) FindOneById(id string) (*User, error) {
	var user User

	query := "SELECT * FROM public.user WHERE id = $1"
	err := r.ql.QueryRowx(query, id).StructScan(&user)
	if err != nil {
		customError := errors.Wrap(err, "readRepository: FindOneById")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return &user, err
}
