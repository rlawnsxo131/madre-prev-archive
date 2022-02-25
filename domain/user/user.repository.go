package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type UserRepository interface {
	FindOneById(id string) (User, error)
}

type repository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindOneById(id string) (User, error) {
	var user User

	sql := "SELECT * FROM user WHERE id = ?"
	err := r.db.QueryRowx(sql, id).StructScan(&user)
	if err != nil {
		err = errors.Wrap(err, "UserRepository: FindOneById sql error")
	}

	return user, err
}
