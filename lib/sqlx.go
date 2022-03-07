package lib

import "database/sql"

var s *sqlxManager

type SqlxManager interface {
	ErrNoRowsReturnRawError(err error, customError error) error
}

type sqlxManager struct{}

func GetSqlxManager() SqlxManager {
	once.Do(func() {
		s = &sqlxManager{}
	})
	return s
}

func (s *sqlxManager) ErrNoRowsReturnRawError(err error, customError error) error {
	if err == sql.ErrNoRows {
		return err
	}
	return customError
}
