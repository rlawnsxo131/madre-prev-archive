package data

import (
	"github.com/pkg/errors"

	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type ReadRepository interface {
	FindAll(limit int) ([]*Data, error)
	FindOneById(id string) (*Data, error)
}

type readRepository struct {
	db     database.Database
	mapper entityMapper
}

func NewReadRepository(db database.Database) ReadRepository {
	return &readRepository{
		db:     db,
		mapper: entityMapper{},
	}
}

func (r *readRepository) FindAll(limit int) ([]*Data, error) {
	var dd []*Data

	query := "SELECT * FROM data" +
		" LIMIT $1"

	rows, err := r.db.Queryx(query, limit)
	if err != nil {
		customError := errors.Wrap(err, "data ReadRepository FindAll query error")
		return nil, utils.ErrNoRowsReturnRawError(err, customError)
	}

	for rows.Next() {
		var d Data
		err := rows.StructScan(&d)
		if err != nil {
			return nil, errors.Wrap(err, "data ReadRepository FindAll StructScan error")
		}
		dd = append(dd, r.mapper.toEntity(&d))
	}

	if err := rows.Close(); err != nil {
		return nil, errors.Wrap(err, "data ReadRepository FindAll rows.Close error")
	}

	return dd, nil
}

func (r *readRepository) FindOneById(id string) (*Data, error) {
	var d Data

	query := "SELECT * FROM data" +
		" WHERE id = $1"

	err := r.db.QueryRowx(query, id).StructScan(&d)
	if err != nil {
		customError := errors.Wrap(err, "DataRepository: FindOneById error")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return r.mapper.toEntity(&d), err
}
