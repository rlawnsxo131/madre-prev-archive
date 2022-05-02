package user

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type ReadRepository interface {
	FindOneById(id string) (*User, error)
}

type readRepository struct {
	db database.Database
}

func NewReadRepository(db database.Database) ReadRepository {
	return &readRepository{
		db: db,
	}
}

func (r *readRepository) FindOneById(id string) (*User, error) {
	var user User

	query := "SELECT * FROM public.user WHERE id = $1"
	err := r.db.QueryRowx(query, id).StructScan(&user)
	if err != nil {
		customError := errors.Wrap(err, "readRepository: FindOneById")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return &user, err
}
