package common

import (
	"database/sql"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/lib/utils"
)

var (
	ErrMissingRequiredValue = errors.New("missing required value")
	ErrNotSupportValue      = errors.New("not support value")
	ErrConflictUniqValue    = errors.New("conflict uniq value")
	ErrUnprocessableValue   = errors.New("unprocessable value")
)

type MadreError struct {
	Err     error
	Code    int
	Message string
}

func NewMadreError(err error, message ...string) *MadreError {
	return &MadreError{
		Err:     err,
		Code:    getHttpStatusCodeFor(err),
		Message: utils.ParseOptionalString(message...),
	}
}

func getHttpStatusCodeFor(err error) int {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return http.StatusNotFound

	case errors.Is(err, pgx.ErrNoRows):
		return http.StatusNotFound

	case errors.Is(err, ErrMissingRequiredValue):
		return http.StatusBadRequest

	case errors.Is(err, ErrNotSupportValue):
		return http.StatusBadRequest

	case errors.Is(err, ErrConflictUniqValue):
		return http.StatusConflict

	case errors.Is(err, ErrUnprocessableValue):
		return http.StatusUnprocessableEntity

	default:
		return http.StatusInternalServerError
	}
}
