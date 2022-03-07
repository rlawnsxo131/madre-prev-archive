package data

import "github.com/jmoiron/sqlx"

type dataWriteRepository struct {
	db *sqlx.DB
}

func NewDataWriteRepository(db *sqlx.DB) *dataWriteRepository {
	return &dataWriteRepository{
		db: db,
	}
}
