package common

import "github.com/pkg/errors"

var (
	ErrMissingRequiredValue = errors.New("missing required value")
	ErrNotSupportValue      = errors.New("not support value")
	ErrConflictUniqValue    = errors.New("conflict uniq value")
	ErrUnProcessableValue   = errors.New("unprocessable value")
)
