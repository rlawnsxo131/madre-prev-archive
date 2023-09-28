package errz

import (
	"fmt"
)

// 필수값이 없을때 발생시킬 에러
type errMissingRequiredValue struct {
	value any
}

func NewErrMissingRequiredValue(value any) error {
	return &errMissingRequiredValue{value}
}

func (e *errMissingRequiredValue) Error() string {
	return fmt.Sprintf("missing required value: %+v", e.value)
}

func IsErrMissingRequiredValue(err error) bool {
	_, ok := err.(*errMissingRequiredValue)
	return ok
}
