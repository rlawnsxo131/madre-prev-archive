package data

import (
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

type DataRepository interface {
	FindAll(limit int) ([]Data, error)
	FindOneById(id string) (Data, error)
}

func NewDataRepository(db *sqlx.DB) DataRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindAll(limit int) ([]Data, error) {
	var dataList []Data
	rows, err := r.db.Queryx("SELECT * FROM data Limit ?", limit)
	for rows.Next() {
		var d Data
		err = rows.StructScan(&d)
		dataList = append(dataList, d)
	}
	return dataList, err
}

func (r *repository) FindOneById(id string) (Data, error) {
	var data Data
	err := r.db.QueryRowx("SELECT * FROM data WHERE id = ?", id).StructScan(&data)
	return data, err
}
