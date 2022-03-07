package data

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib"
)

var sqlxManager = lib.NewSqlxManager()

type dataReadRepository struct {
	db *sqlx.DB
}

func NewDataReadRepository(db *sqlx.DB) *dataReadRepository {
	return &dataReadRepository{
		db: db,
	}
}

func (r *dataReadRepository) FindAll(limit int) ([]Data, error) {
	var dataList []Data

	query := "SELECT * FROM data Limit ?"
	rows, err := r.db.Queryx(query, limit)
	if err != nil {
		customError := errors.Wrap(err, "DataRepository: FindAll query error")
		return nil, sqlxManager.ErrNoRowsReturnRawError(err, customError)
	}

	for rows.Next() {
		var d Data
		err := rows.StructScan(&d)
		if err != nil {
			return nil, errors.Wrap(err, "DataRepository: FindAll StructScan error")
		}
		dataList = append(dataList, d)
	}

	return dataList, nil
}

func (r *dataReadRepository) FindOneById(id int64) (Data, error) {
	var data Data

	query := "SELECT * FROM data WHERE id = ?"
	err := r.db.QueryRowx(query, id).StructScan(&data)
	if err != nil {
		customError := errors.Wrap(err, "DataRepository: FindOneById error")
		err = sqlxManager.ErrNoRowsReturnRawError(err, customError)
	}

	return data, err
}

func (r *dataReadRepository) FindOneByUUID(uuid string) (Data, error) {
	var data Data

	query := "SELECT * FROM data WHERE uuid = ?"
	err := r.db.QueryRowx(query, uuid).StructScan(&data)
	if err != nil {
		if err != sql.ErrNoRows {
			customError := errors.Wrap(err, "DataRepository: FindOneByUUID error")
			err = sqlxManager.ErrNoRowsReturnRawError(err, customError)
		}
	}

	return data, err
}
