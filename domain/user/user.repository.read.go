package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type UserReadRepository interface {
	FindOneById(id int64) (User, error)
	FindOneByUUID(uuid string) (User, error)
}

type userReadRepository struct {
	ql logger.QueryLogger
}

func NewUserReadRepository(db *sqlx.DB) UserReadRepository {
	return &userReadRepository{
		ql: logger.NewQueryLogger(db),
	}
}

func (r *userReadRepository) FindOneById(id int64) (User, error) {
	var user User

	query := "SELECT * FROM user WHERE id = ?"
	err := r.ql.QueryRowx(query, id).StructScan(&user)
	if err != nil {
		customError := errors.Wrap(err, "UserRepository: FindOneById")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return user, err
}

func (r *userReadRepository) FindOneByUUID(uuid string) (User, error) {
	var user User

	query := "SELECT * FROM user WHERE uuid = ?"
	err := r.ql.QueryRowx(query, uuid).StructScan(&user)
	if err != nil {
		customError := errors.Wrap(err, "UserRepository: FindOneByUUID")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return user, err
}
