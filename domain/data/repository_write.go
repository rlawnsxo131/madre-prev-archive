package data

import (
	"github.com/rlawnsxo131/madre-server-v2/database"
)

type WriteRepository interface{}

type writeRepository struct {
	db     database.Database
	mapper entityMapper
}

func NewWriteRepository(db database.Database) WriteRepository {
	return &writeRepository{
		db:     db,
		mapper: entityMapper{},
	}
}
