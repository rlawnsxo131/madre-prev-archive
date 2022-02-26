package data

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type DataRepository interface {
	FindAll(limit int) ([]Data, error)
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

func (r *repository) FindAll(limit int) ([]Data, error) {
	var dataList []Data

	sql := "SELECT * FROM data Limit ?"
	rows, err := r.db.Queryx(sql, limit)
	if err != nil {
		return nil, errors.Wrap(err, "DataRepository: FindAll sql error")
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

func (r *repository) FindOneById(id string) (Data, error) {
	var data Data

	sql := "SELECT * FROM data WHERE id = ?"
	err := r.db.QueryRowx(sql, id).StructScan(&data)
	if err != nil {
		err = errors.Wrap(err, "DataRepository: FindOneById error")
	}

	return data, err
}
