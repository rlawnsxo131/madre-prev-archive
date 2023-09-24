package domainerr

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"
)

type DomainError struct {
	Err  error
	Code int
	Msg  string
}

func New(err error, message ...string) *DomainError {
	return &DomainError{
		Err:  err,
		Code: getHttpStatusCodeFor(err),
		Msg:  strings.Join(message, ""),
	}
}

func getHttpStatusCodeFor(err error) int {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return http.StatusNotFound

	case IsErrMissingRequiredValue(err):
		return http.StatusBadRequest

	case IsErrNotSupportValue(err):
		return http.StatusBadRequest

	case IsErrConflictUniqValue(err):
		return http.StatusConflict

	case IsErrUnprocessableValue(err):
		return http.StatusUnprocessableEntity

	default:
		return http.StatusInternalServerError
	}
}
