package account

import (
	"github.com/pkg/errors"
)

var (
	ErrInvalidUsername       = errors.New("username is required")
	ErrInvalidEmail          = errors.New("email is required")
	ErrInvalidSocialId       = errors.New("socialId is required")
	ErrInvalidProvider       = errors.New("provider is required")
	ErrInvalidProviderGoogle = errors.New("provider is not google")
)
