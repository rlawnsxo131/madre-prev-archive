package errz

import "fmt"

// 지원하지 않는 값이나 포멧일때 발생시킬 에러
type errNotSupportValue struct {
	value any
}

func NewErrNotSupportValue(value any) error {
	return &errNotSupportValue{value}
}

func (e *errNotSupportValue) Error() string {
	return fmt.Sprintf("not support value: %+v", e.value)
}

func IsErrNotSupportValue(err error) bool {
	_, ok := err.(*errNotSupportValue)
	return ok
}
