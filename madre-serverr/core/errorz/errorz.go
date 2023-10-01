package errorz

import (
	"runtime/debug"
	"strings"
)

type withStack struct {
	err   error
	stack string
}

func New(err error) error {
	return &withStack{
		err:   err,
		stack: string(debug.Stack()),
	}
}

func (s *withStack) Error() string {
	return strings.Join(
		[]string{
			"Error: ",
			s.err.Error(),
			", Stack: ",
			s.stack,
		},
		"",
	)
}
