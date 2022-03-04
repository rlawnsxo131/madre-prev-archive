package user

import "github.com/jmoiron/sqlx"

type UserWriteRepository interface{}

type userWriteRepository struct {
	db *sqlx.DB
}

func NewUserWriteRepository(db *sqlx.DB) UserWriteRepository {
	return &userWriteRepository{
		db: db,
	}
}
