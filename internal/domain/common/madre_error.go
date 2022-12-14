package common

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/utils"
)

var (
	ErrMissingRequiredValue = errors.New("missing required value")
	ErrNotSupportValue      = errors.New("not support value")
	ErrConflictUniqValue    = errors.New("conflict uniq value")
	ErrUnprocessableValue   = errors.New("unprocessable value")
)

type MadreError struct {
	Err     error
	Message string
}

func NewMadreError(err error, message ...string) *MadreError {
	return &MadreError{
		Err:     err,
		Message: utils.ParseOtionalString(message...),
	}
}
