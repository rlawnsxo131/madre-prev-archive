package lib

import "database/sql"

type SqlxManager interface {
	ErrNoRowsReturnRawError(err error, customError error) error
}

type sqlxManager struct{}

var sm *sqlxManager

func NewSqlxManager() SqlxManager {
	if sm == nil {
		sm = &sqlxManager{}
	}
	return sm
}

func (s *sqlxManager) ErrNoRowsReturnRawError(err error, customError error) error {
	if err == sql.ErrNoRows {
		return err
	}
	return customError
}
