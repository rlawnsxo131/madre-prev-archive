package errz

import (
	"database/sql"
	"errors"
	"net/http"
)

// @TODO 에러 처리 어떻게할지 고민중

func GetHttpStatusCodeFor(err error) int {
	// validationErrors, ok := err.(validator.ValidationErrors)

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
