package utils

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
)

// default values
func Contains[T int | string](ss []T, value T) bool {
	for _, s := range ss {
		if s == value {
			return true
		}
	}
	return false
}

func IfIsNotExistGetDefaultIntValue(value int, defaultValue int) int {
	if value == 0 {
		value = defaultValue
	}
	return value
}

func ParseOptionalString(ss ...string) string {
	var result string
	if len(ss) > 0 {
		for _, v := range ss {
			result += v
		}
	}
	return result
}

// database values
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
	if err == pgx.ErrNoRows {
		return err
	}
	return customError
}
