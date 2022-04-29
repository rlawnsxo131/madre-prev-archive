package data

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type ReadRepository interface {
	FindAll(limit int) ([]*Data, error)
	FindOneById(id string) (*Data, error)
}

type readRepository struct {
	ql logger.QueryLogger
}

func NewReadRepository(db *sqlx.DB) ReadRepository {
	return &readRepository{
		ql: logger.NewQueryLogger(db),
	}
}

func (r *readRepository) FindAll(limit int) ([]*Data, error) {
	var dataList []*Data

	query := "SELECT * FROM data LIMIT $1"
	rows, err := r.ql.Queryx(query, limit)
	if err != nil {
		customError := errors.Wrap(err, "DataRepository: FindAll query error")
		return nil, utils.ErrNoRowsReturnRawError(err, customError)
	}

	for rows.Next() {
		var d Data
		err := rows.StructScan(&d)
		if err != nil {
			return nil, errors.Wrap(err, "DataRepository: FindAll StructScan error")
		}
		dataList = append(dataList, &d)
	}

	if err := rows.Close(); err != nil {
		return nil, errors.Wrap(err, "DataRepository: FindAll rows.Close error")
	}

	return dataList, nil
}

func (r *readRepository) FindOneById(id string) (*Data, error) {
	var data Data

	query := "SELECT * FROM data WHERE id = $1"
	err := r.ql.QueryRowx(query, id).StructScan(&data)
	if err != nil {
		customError := errors.Wrap(err, "DataRepository: FindOneById error")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return &data, err
}
