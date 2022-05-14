package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type singletonDatabase struct {
	DB *sqlx.DB
	l  *zerolog.Logger
}

func (sd *singletonDatabase) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	sd.l.Log().Timestamp().Str("query", fmt.Sprintf("%s,%+v", query, args)).Msg("")
	return sd.DB.Queryx(query, args...)
}

func (sd *singletonDatabase) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	sd.l.Log().Timestamp().Str("query", fmt.Sprintf("%s,%+v", query, args)).Msg("")
	return sd.DB.QueryRowx(query, args...)
}

func (sd *singletonDatabase) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	sd.l.Log().Timestamp().Str("query", fmt.Sprintf("%s,%+v", query, arg)).Msg("")
	return sd.DB.NamedQuery(query, arg)
}

func (sd *singletonDatabase) PrepareNamedGet(id *string, query string, arg interface{}) error {
	sd.l.Log().Timestamp().Str("query", fmt.Sprintf("%s,%+v", query, arg)).Msg("")
	stmt, err := sd.DB.PrepareNamed(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	return stmt.Get(id, arg)
}
