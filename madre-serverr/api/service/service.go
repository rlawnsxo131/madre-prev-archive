package service

import (
	"strings"
)

type ErrWithHTTPCode struct {
	Err     error
	Code    int
	Message string
}

func NewErrWithHTTPCode(err error, code int, message ...string) *ErrWithHTTPCode {
	return &ErrWithHTTPCode{
		Err:     err,
		Code:    code,
		Message: strings.Join(message, ""),
	}
}
