package account

import (
	"regexp"

	"github.com/pkg/errors"
)

var (
	ErrInvalidUsernameFormat = errors.New("username regex MatchString error")
)

type Username struct {
	name string
}

func NewUsername(name string) (*Username, error) {
	match, err := regexp.MatchString(
		"^[a-zA-Z0-9]{1,20}$",
		name,
	)
	if err != nil {
		return nil, errors.Wrap(err, "username regex MatchString error")
	}
	if !match {
		return nil, ErrInvalidUsernameFormat
	}

	return &Username{
		name: name,
	}, nil
}
