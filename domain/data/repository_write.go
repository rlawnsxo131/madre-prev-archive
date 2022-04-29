package data

import "github.com/jmoiron/sqlx"

type DataWriteRepository interface{}

type dataWriteRepository struct {
	db *sqlx.DB
}

func NewDataWriteRepository(db *sqlx.DB) DataWriteRepository {
	return &dataWriteRepository{
		db: db,
	}
}
