package errz

import "fmt"

// 처리할수 없는 값에대한 에러
// json, regex 등을 파싱하는 자체에서 에러가 발생하는 경우 사용
type errUnprocessableValue struct {
	value any
}

func NewErrUnprocessableValue(value any) error {
	return &errUnprocessableValue{value}
}

func (e *errUnprocessableValue) Error() string {
	return fmt.Sprintf("unprocessable value: %+v", e.value)
}

func IsErrUnprocessableValue(err error) bool {
	_, ok := err.(*errUnprocessableValue)
	return ok
}
