package errz

import (
	"database/sql"
	"errors"
	"net/http"
)

func GetHttpStatusCodeFor(err error) int {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return http.StatusNotFound

	case IsErrMissingRequiredValue(err):
		return http.StatusBadRequest

	case IsErrNotSupportValue(err) || IsErrUnprocessableValue(err):
		return http.StatusUnprocessableEntity

	case IsErrConflictUniqValue(err):
		return http.StatusConflict

	default:
		return http.StatusInternalServerError
	}
}
