package user

import (
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

type UserRepository interface {
	FindOneById(id string) (User, error)
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindOneById(id string) (User, error) {
	var user User
	err := r.db.QueryRowx("SELECT * FROM user WHERE id = ?", id).StructScan(&user)
	return user, err
}
