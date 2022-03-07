package utils

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

var ValidateManager *validator.Validate = validator.New()

func GenerateUUIDString() string {
	return uuid.NewV4().String()
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
