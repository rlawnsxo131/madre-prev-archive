package errz

import (
	"fmt"
)

// 유니크한 값이 중복되었을때 내보내는 에러
type errConflictUniqValue struct {
	value any
}

func NewErrConflictUniqValue(value any) error {
	return &errConflictUniqValue{value}
}

func (e *errConflictUniqValue) Error() string {
	return fmt.Sprintf("conflict unique value: %+v", e.value)
}

func IsErrConflictUniqValue(err error) bool {
	_, ok := err.(*errConflictUniqValue)
	return ok
}
