package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type AuthRepository interface{}

type repository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindOneById(key string) (Auth, error) {
	var auth Auth

	sql := "SELECT * FROM auth WHERE id = ?"
	err := r.db.QueryRowx(sql, key).StructScan(&auth)
	if err != nil {
		err = errors.Wrap(err, "AuthRepository: FindOneById error")
	}

	return auth, err
}
