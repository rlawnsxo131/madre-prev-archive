package user

import "github.com/jmoiron/sqlx"

type userWriteRepository struct {
	db *sqlx.DB
}

func NewUserWriteRepository(db *sqlx.DB) *userWriteRepository {
	return &userWriteRepository{
		db: db,
	}
}
