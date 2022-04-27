package utils

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var Validator *validator.Validate = validator.New()

func GenerateUUIDString() string {
	return uuid.NewString()
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func NormalizeNullString(sn sql.NullString) string {
	if sn.Valid {
		return sn.String
	}
	return ""
}

func ErrNoRowsReturnRawError(err error, customError error) error {
	if err == sql.ErrNoRows {
		return err
	}
	return customError
}

func IfIsNotExistGetDefaultIntValue(value int, defaultValue int) int {
	if value == 0 {
		value = defaultValue
	}
	return value
}
