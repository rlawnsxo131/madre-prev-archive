package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib"
)

var sqlxManager = lib.NewSqlxManager()

type userReadRepository struct {
	db *sqlx.DB
}

func NewUserReadRepository(db *sqlx.DB) *userReadRepository {
	return &userReadRepository{
		db: db,
	}
}

func (r *userReadRepository) FindOneById(id int64) (User, error) {
	var user User

	query := "SELECT * FROM user WHERE id = ?"
	err := r.db.QueryRowx(query, id).StructScan(&user)
	if err != nil {
		customError := errors.Wrap(err, "UserRepository: FindOneById error")
		err = sqlxManager.ErrNoRowsReturnRawError(err, customError)
	}

	return user, err
}

func (r *userReadRepository) FindOneByUUID(uuid string) (User, error) {
	var user User

	query := "SELECT * FROM user WHERE uuid = ?"
	err := r.db.QueryRowx(query, uuid).StructScan(&user)
	if err != nil {
		customError := errors.Wrap(err, "UserRepository: FindOneByUUID error")
		err = sqlxManager.ErrNoRowsReturnRawError(err, customError)
	}

	return user, err
}
