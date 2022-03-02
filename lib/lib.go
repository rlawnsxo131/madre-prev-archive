package lib

import (
	uuid "github.com/satori/go.uuid"
)

/**
 * uuid
 * github.com/satori/go.uuid
 */
func GenerateUUID() string {
	uuid := uuid.NewV4()
	return uuid.String()
}

/**
 * default value
 */
func IfIsNotExistGetDefaultIntValue(value int, defaultValue int) int {
	if value == 0 {
		value = defaultValue
	}
	return value
}
