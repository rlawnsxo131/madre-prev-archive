package data

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type DataRepository interface {
	FindAll(limit int) ([]Data, error)
	FindOneById(id uint) (Data, error)
	FindOneByUUID(uuid string) (Data, error)
}

type dataRepository struct {
	db *sqlx.DB
}

func NewDataRepository(db *sqlx.DB) DataRepository {
	return &dataRepository{
		db: db,
	}
}

func (r *dataRepository) FindAll(limit int) ([]Data, error) {
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

func (r *dataRepository) FindOneById(id uint) (Data, error) {
	var data Data

	sql := "SELECT * FROM data WHERE id = ?"
	err := r.db.QueryRowx(sql, id).StructScan(&data)
	if err != nil {
		err = errors.Wrap(err, "DataRepository: FindOneById error")
	}

	return data, err
}

func (r *dataRepository) FindOneByUUID(uuid string) (Data, error) {
	var data Data

	sql := "SELECT * FROM data WHERE uuid = ?"
	err := r.db.QueryRowx(sql, uuid).StructScan(&data)
	if err != nil {
		err = errors.Wrap(err, "DataRepository: FindOneByUUID error")
	}

	return data, err
}
