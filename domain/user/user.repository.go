package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type UserRepository interface {
	FindOneById(id uint) (User, error)
	FindOneByUUID(uuid string) (User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindOneById(id uint) (User, error) {
	var user User

	sql := "SELECT * FROM user WHERE id = ?"
	err := r.db.QueryRowx(sql, id).StructScan(&user)
	if err != nil {
		err = errors.Wrap(err, "UserRepository: FindOneById sql error")
	}

	return user, err
}

func (r *userRepository) FindOneByUUID(uuid string) (User, error) {
	var user User

	sql := "SELECT * FROM user WHERE uuid = ?"
	err := r.db.QueryRowx(sql, uuid).StructScan(&user)
	if err != nil {
		err = errors.Wrap(err, "UserRepository: FindOneByUUID sql error")
	}

	return user, err
}
