package lib

import "database/sql"

type SqlxLib interface {
	ErrNoRowsReturnRawError(err error, customError error) error
}

type sqlxLib struct{}

func NewSqlxLib() SqlxLib {
	return &sqlxLib{}
}

func (s *sqlxLib) ErrNoRowsReturnRawError(err error, customError error) error {
	if err == sql.ErrNoRows {
		return err
	}
	return customError
}
