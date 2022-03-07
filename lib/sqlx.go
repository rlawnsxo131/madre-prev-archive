package lib

import "database/sql"

type sqlxManager struct{}

func NewSqlxManager() *sqlxManager {
	return &sqlxManager{}
}

func (s *sqlxManager) ErrNoRowsReturnRawError(err error, customError error) error {
	if err == sql.ErrNoRows {
		return err
	}
	return customError
}
