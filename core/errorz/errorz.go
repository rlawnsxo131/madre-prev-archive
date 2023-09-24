package errorz

import (
	"runtime/debug"
	"strings"
)

/* @TODO stack 관련 고민 해보기 */
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
