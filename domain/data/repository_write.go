package data

import "github.com/jmoiron/sqlx"

type WriteRepository interface{}

type writeRepository struct {
	db *sqlx.DB
}

func NewWriteRepository(db *sqlx.DB) WriteRepository {
	return &writeRepository{
		db: db,
	}
}
