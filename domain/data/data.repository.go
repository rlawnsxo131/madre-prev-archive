package data

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type DataRepository interface {
	FindAll(limit int) (list []Data, err error)
	FindOneById(id string) (Data, error)
}

type repository struct {
	db *sqlx.DB
}

func NewDataRepository(db *sqlx.DB) DataRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindAll(limit int) (list []Data, err error) {
	var dataList []Data

	sql := "SELECT * FROM data Limit ?"
	rows, err := r.db.Queryx(sql, limit)
	if err != nil {
		err = errors.Wrap(err, "DataRepository: FindAll sql error")
		return nil, err
	}

	for rows.Next() {
		var d Data
		err := rows.StructScan(&d)
		if err != nil {
			err = errors.Wrap(err, "DataRepository: FindAll StructScan error")
			return nil, err
		}
		dataList = append(dataList, d)
	}

	return dataList, nil
}

func (r *repository) FindOneById(id string) (Data, error) {
	var data Data

	sql := "SELECT * FROM data WHERE id = ?"
	err := r.db.QueryRowx(sql, id).StructScan(&data)
	if err != nil {
		err = errors.Wrap(err, "DataRepository: FindOneById error")
	}

	return data, err
}
