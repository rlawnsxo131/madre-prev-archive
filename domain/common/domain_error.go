package common

import (
	"errors"
)

var (
	ErrMissingRequiredValue = errors.New("missing required value")
	ErrNotSupportValue      = errors.New("not support value")
	ErrConflictUniqValue    = errors.New("conflict unique value")
	ErrUnprocessableValue   = errors.New("unprocessable value")
)
